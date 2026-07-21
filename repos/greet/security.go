package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type SecurityInfo struct {
	FailedCount  int
	LastAttacker string
	LastLogin    string
}

func getSecurityInfo() SecurityInfo {
	return SecurityInfo{
		FailedCount:  getFailedSSHCount(),
		LastAttacker: getLastAttacker(),
		LastLogin:    getLastLogin(),
	}
}

func getFailedSSHCount() int {
	out, err := exec.Command("sh", "-c",
		`journalctl --since today -q --no-pager 2>/dev/null | grep -cE "Failed password|Invalid user" 2>/dev/null || echo 0`,
	).Output()
	if err != nil {
		return 0
	}
	s := strings.TrimSpace(string(out))
	n := 0
	fmt.Sscanf(s, "%d", &n)
	return n
}

func getLastAttacker() string {
	out, err := exec.Command("sh", "-c",
		`journalctl --since today -q --no-pager 2>/dev/null | grep -E "Failed password|Invalid user" | tail -1`,
	).Output()
	if err != nil || len(strings.TrimSpace(string(out))) == 0 {
		return "none today"
	}
	line := string(out)
	if idx := strings.Index(line, "from "); idx != -1 {
		rest := line[idx+5:]
		parts := strings.Fields(rest)
		if len(parts) > 0 {
			return parts[0]
		}
	}
	return "none today"
}

func getLastLogin() string {
	out, err := exec.Command("last", "-n", "3", "-F", "vincent").Output()
	if err != nil || len(out) == 0 {
		return "unknown"
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, line := range lines {
		if strings.Contains(line, "still logged in") {
			continue
		}
		if !strings.HasPrefix(line, "vincent") {
			continue
		}
		fields := strings.Fields(line)
		// Format: vincent pts/0 IP Weekday Mon DD HH:MM:SS YYYY
		if len(fields) >= 9 {
			timeStr := fmt.Sprintf("%s %s %s %s %s",
				fields[4], fields[5], fields[6], fields[7], fields[8])
			t, err := time.Parse("Mon Jan 2 15:04:05 2006", timeStr)
			if err == nil {
				return "vincent  ·  " + formatAgo(time.Since(t))
			}
		}
		return "vincent"
	}
	return "vincent"
}

func formatAgo(d time.Duration) string {
	if d < time.Minute {
		return "just now"
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm ago", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("%dh ago", int(d.Hours()))
	}
	return fmt.Sprintf("%dd ago", int(d.Hours())/24)
}