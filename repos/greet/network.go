package main

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"
)

type NetworkInfo struct {
	Interface string
	Address   string
	Tunnel    string
	Internet  string
}

func getNetworkInfo() NetworkInfo {
	iface, addr := getLocalIP()
	return NetworkInfo{
		Interface: iface,
		Address:   addr,
		Tunnel:    getTunnelStatus(),
		Internet:  getInternetLatency(),
	}
}

func getLocalIP() (iface string, addr string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "unknown", "unknown"
	}

	// Prefer wlan0, eth0, enp*, wlp* — skip lo and docker/veth
	preferred := []string{"eth0", "wlan0"}
	skip := []string{"lo", "docker", "veth", "br-", "virbr"}

	candidates := make(map[string]string)
	for _, ifc := range interfaces {
		name := ifc.Name

		shouldSkip := false
		for _, s := range skip {
			if strings.HasPrefix(name, s) {
				shouldSkip = true
				break
			}
		}
		if shouldSkip {
			continue
		}

		addrs, err := ifc.Addrs()
		if err != nil {
			continue
		}
		for _, a := range addrs {
			var ip net.IP
			switch v := a.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() || ip.To4() == nil {
				continue
			}
			candidates[name] = ip.String()
		}
	}

	// Return preferred if found
	for _, p := range preferred {
		if addr, ok := candidates[p]; ok {
			return p, addr
		}
	}
	// Fall back to first candidate
	for name, addr := range candidates {
		return name, addr
	}
	return "unknown", "unknown"
}

func getTunnelStatus() string {
	// Check if cloudflared.service is active via systemctl
	out, err := exec.Command("systemctl", "is-active", "cloudflared").Output()
	if err != nil {
		return "inactive"
	}
	state := strings.TrimSpace(string(out))
	if state == "active" {
		return "active"
	}
	return state
}

func getInternetLatency() string {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", "1.1.1.1:53", 3*time.Second)
	if err != nil {
		return "unreachable"
	}
	conn.Close()
	ms := time.Since(start).Milliseconds()
	return fmt.Sprintf("%d ms", ms)
}