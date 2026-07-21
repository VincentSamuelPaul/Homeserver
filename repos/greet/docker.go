package main

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"strings"
	"time"
)

type ContainerInfo struct {
	Name   string
	State  string
	Uptime string
}

// displayName → actual container name patterns to match
var serviceMap = []struct {
	display string
	match   []string // substrings to match against container name
}{
	{"portfolio", []string{"portfolio"}},
	{"nextcloud", []string{"nextcloud-app", "nextcloud"}},
	{"nextcloud-db", []string{"nextcloud-db"}},
	{"nextcloud-redis", []string{"nextcloud-redis"}},
	{"nextcloud-nginx", []string{"nextcloud-nginx"}},
	{"grafana", []string{"grafana"}},
	{"loki", []string{"loki"}},
	{"fluent-bit", []string{"fluent-bit"}},
	{"whatsapp-bot", []string{"whatsapp-bot"}},
	{"docs", []string{"docs"}},
}

type dockerContainer struct {
	Names  []string `json:"Names"`
	State  string   `json:"State"`
	Status string   `json:"Status"`
}

func getContainers() []ContainerInfo {
	client := &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/var/run/docker.sock")
			},
		},
	}

	resp, err := client.Get("http://localhost/containers/json?all=1")
	if err != nil {
		return fallbackContainers()
	}
	defer resp.Body.Close()

	var raw []dockerContainer
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return fallbackContainers()
	}

	// Build raw name → container map (strip leading /)
	byName := make(map[string]dockerContainer)
	for _, c := range raw {
		if len(c.Names) > 0 {
			n := strings.TrimPrefix(c.Names[0], "/")
			byName[n] = c
		}
	}

	result := make([]ContainerInfo, 0, len(serviceMap))
	for _, svc := range serviceMap {
		found := false
		for _, pattern := range svc.match {
			for rawName, c := range byName {
				if strings.Contains(rawName, pattern) {
					result = append(result, ContainerInfo{
						Name:   svc.display,
						State:  c.State,
						Uptime: parseUptime(c.Status),
					})
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			result = append(result, ContainerInfo{Name: svc.display, State: "stopped"})
		}
	}
	return result
}

func parseUptime(status string) string {
	status = strings.TrimSpace(status)
	if !strings.HasPrefix(status, "Up ") {
		return ""
	}
	s := strings.TrimPrefix(status, "Up ")
	s = strings.TrimSpace(s)
	replacer := strings.NewReplacer(
		"About an hour", "~1h",
		"About a minute", "~1m",
		" hours", "h",
		" hour", "h",
		" days", "d",
		" day", "d",
		" minutes", "m",
		" minute", "m",
		" seconds", "s",
		" second", "s",
		" weeks", "w",
		" week", "w",
	)
	return replacer.Replace(s)
}

func fallbackContainers() []ContainerInfo {
	result := make([]ContainerInfo, len(serviceMap))
	for i, svc := range serviceMap {
		result[i] = ContainerInfo{Name: svc.display, State: "unknown"}
	}
	return result
}