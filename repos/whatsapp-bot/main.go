package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// ── config ────────────────────────────────────────────────────────────────────

type Config struct {
	WhatsAppToken       string
	WhatsAppPhoneID     string
	WhatsAppVerifyToken string
	NextcloudURL        string
	NextcloudUser       string
	NextcloudPassword   string
}

func loadConfig() Config {
	return Config{
		WhatsAppToken:       os.Getenv("WHATSAPP_TOKEN"),
		WhatsAppPhoneID:     os.Getenv("WHATSAPP_PHONE_ID"),
		WhatsAppVerifyToken: os.Getenv("WHATSAPP_VERIFY_TOKEN"),
		NextcloudURL:        os.Getenv("NEXTCLOUD_URL"),
		NextcloudUser:       os.Getenv("NEXTCLOUD_USER"),
		NextcloudPassword:   os.Getenv("NEXTCLOUD_PASSWORD"),
	}
}

// ── whatsapp types ────────────────────────────────────────────────────────────

type WebhookPayload struct {
	Entry []struct {
		Changes []struct {
			Value struct {
				Messages []Message `json:"messages"`
			} `json:"value"`
		} `json:"changes"`
	} `json:"entry"`
}

type Message struct {
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Image    *Media   `json:"image"`
	Video    *Media   `json:"video"`
	Audio    *Media   `json:"audio"`
	Document *Media   `json:"document"`
}

type Media struct {
	ID       string `json:"id"`
	MimeType string `json:"mime_type"`
	Filename string `json:"filename"`
}

type MediaInfo struct {
	URL      string `json:"url"`
	MimeType string `json:"mime_type"`
}

// ── helpers ───────────────────────────────────────────────────────────────────

func getFileType(mimeType string) string {
	switch {
	case strings.HasPrefix(mimeType, "image/"):
		return "images"
	case strings.HasPrefix(mimeType, "video/"):
		return "videos"
	case strings.HasPrefix(mimeType, "audio/"):
		return "audio"
	case strings.Contains(mimeType, "pdf"),
		strings.Contains(mimeType, "document"),
		strings.Contains(mimeType, "sheet"),
		strings.Contains(mimeType, "presentation"),
		strings.HasPrefix(mimeType, "text/"):
		return "documents"
	default:
		return "other"
	}
}

func getExtension(mimeType string) string {
	parts := strings.Split(mimeType, "/")
	if len(parts) == 2 {
		ext := parts[1]
		ext = strings.Split(ext, ";")[0]
		return ext
	}
	return "bin"
}

func dateFolder() string {
	return time.Now().Format("2006-01-02")
}

func basicAuth(user, pass string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+pass))
}

// ── whatsapp api ──────────────────────────────────────────────────────────────

func (c Config) getMediaInfo(mediaID string) (*MediaInfo, error) {
	req, _ := http.NewRequest("GET",
		fmt.Sprintf("https://graph.facebook.com/v19.0/%s", mediaID), nil)
	req.Header.Set("Authorization", "Bearer "+c.WhatsAppToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var info MediaInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}

func (c Config) downloadMedia(mediaURL string) ([]byte, error) {
	req, _ := http.NewRequest("GET", mediaURL, nil)
	req.Header.Set("Authorization", "Bearer "+c.WhatsAppToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (c Config) sendReply(to, message string) error {
	body := map[string]any{
		"messaging_product": "whatsapp",
		"to":                to,
		"type":              "text",
		"text":              map[string]string{"body": message},
	}

	data, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST",
		fmt.Sprintf("https://graph.facebook.com/v19.0/%s/messages", c.WhatsAppPhoneID),
		bytes.NewReader(data))
	req.Header.Set("Authorization", "Bearer "+c.WhatsAppToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// ── nextcloud webdav ──────────────────────────────────────────────────────────

func (c Config) ensureFolder(path string) error {
	parts := strings.Split(path, "/")
	encoded := make([]string, len(parts))
	for i, p := range parts {
		encoded[i] = url.PathEscape(p)
	}

	webdavURL := fmt.Sprintf("%s/remote.php/dav/files/%s/%s",
		c.NextcloudURL, c.NextcloudUser, strings.Join(encoded, "/"))

	req, _ := http.NewRequest("MKCOL", webdavURL, nil)
	req.Header.Set("Authorization", basicAuth(c.NextcloudUser, c.NextcloudPassword))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 405 = folder already exists, that's fine
	if resp.StatusCode != 201 && resp.StatusCode != 405 {
		return fmt.Errorf("MKCOL failed: %d", resp.StatusCode)
	}
	return nil
}

func (c Config) uploadToNextcloud(data []byte, filename, mimeType string) (string, error) {
	fileType := getFileType(mimeType)
	date := dateFolder()

	// ensure folder hierarchy
	folders := []string{
		"WhatsApp Backup",
		fmt.Sprintf("WhatsApp Backup/%s", fileType),
		fmt.Sprintf("WhatsApp Backup/%s/%s", fileType, date),
	}
	for _, f := range folders {
		if err := c.ensureFolder(f); err != nil {
			return "", fmt.Errorf("ensureFolder %s: %w", f, err)
		}
	}

	remotePath := fmt.Sprintf("WhatsApp Backup/%s/%s/%s", fileType, date, filename)
	parts := strings.Split(remotePath, "/")
	encoded := make([]string, len(parts))
	for i, p := range parts {
		encoded[i] = url.PathEscape(p)
	}

	webdavURL := fmt.Sprintf("%s/remote.php/dav/files/%s/%s",
		c.NextcloudURL, c.NextcloudUser, strings.Join(encoded, "/"))

	req, _ := http.NewRequest("PUT", webdavURL, bytes.NewReader(data))
	req.Header.Set("Authorization", basicAuth(c.NextcloudUser, c.NextcloudPassword))
	req.Header.Set("Content-Type", mimeType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 && resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload failed %d: %s", resp.StatusCode, body)
	}

	return remotePath, nil
}

// ── handlers ──────────────────────────────────────────────────────────────────

func (c Config) verifyWebhook(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")

	if mode == "subscribe" && token == c.WhatsAppVerifyToken {
		log.Println("Webhook verified")
		w.WriteHeader(200)
		fmt.Fprint(w, challenge)
		return
	}
	w.WriteHeader(403)
}

func (c Config) handleWebhook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200) // respond fast to Meta

	var payload WebhookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println("decode error:", err)
		return
	}

	if len(payload.Entry) == 0 || len(payload.Entry[0].Changes) == 0 {
		return
	}

	messages := payload.Entry[0].Changes[0].Value.Messages
	if len(messages) == 0 {
		return
	}

	msg := messages[0]
	from := msg.From

	// get media based on type
	var media *Media
	switch msg.Type {
	case "image":
		media = msg.Image
	case "video":
		media = msg.Video
	case "audio":
		media = msg.Audio
	case "document":
		media = msg.Document
	default:
		c.sendReply(from, "Send me any file and I'll back it up to Nextcloud. 📁")
		return
	}

	if media == nil {
		return
	}

	// generate filename if not provided
	filename := media.Filename
	if filename == "" {
		filename = fmt.Sprintf("%s-%d.%s",
			msg.Type, time.Now().UnixMilli(), getExtension(media.MimeType))
	}

	go func() {
		c.sendReply(from, fmt.Sprintf("⏳ Backing up %s...", filename))

		info, err := c.getMediaInfo(media.ID)
		if err != nil {
			log.Println("getMediaInfo error:", err)
			c.sendReply(from, "❌ Failed to get media info.")
			return
		}

		data, err := c.downloadMedia(info.URL)
		if err != nil {
			log.Println("downloadMedia error:", err)
			c.sendReply(from, "❌ Failed to download media.")
			return
		}

		remotePath, err := c.uploadToNextcloud(data, filename, media.MimeType)
		if err != nil {
			log.Println("uploadToNextcloud error:", err)
			c.sendReply(from, "❌ Failed to upload to Nextcloud.")
			return
		}

		c.sendReply(from, fmt.Sprintf("✓ Backed up to:\n%s", remotePath))
		log.Printf("backed up: %s", remotePath)
	}()
}

// ── main ──────────────────────────────────────────────────────────────────────

func main() {
	cfg := loadConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /webhook", cfg.verifyWebhook)
	mux.HandleFunc("POST /webhook", cfg.handleWebhook)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	log.Println("whatsapp backup bot running on :3001")
	log.Fatal(http.ListenAndServe(":3001", mux))
}
