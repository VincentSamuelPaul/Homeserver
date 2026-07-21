package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type SystemInfo struct {
	Host      string
	Uptime    string
	Kernel    string
	CPUPct    float64
	MemUsed   float64
	MemTotal  float64
	RootUsed  float64
	RootTotal float64
	HomeUsed  float64
	HomeTotal float64
}

func getSystemInfo() SystemInfo {
	rootUsed, rootTotal := getDiskStats("/")
	homeUsed, homeTotal := getDiskStats("/home")
	return SystemInfo{
		Host:      getHostname(),
		Uptime:    getUptime(),
		Kernel:    getKernel(),
		CPUPct:    getCPUPercent(),
		MemUsed:   getMemUsed(),
		MemTotal:  getMemTotal(),
		RootUsed:  rootUsed,
		RootTotal: rootTotal,
		HomeUsed:  homeUsed,
		HomeTotal: homeTotal,
	}
}

func getHostname() string {
	b, err := os.ReadFile("/etc/hostname")
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(b))
}

func getUptime() string {
	ist := time.FixedZone("IST", 5*60*60+30*60)
	start := time.Date(2026, 7, 1, 0, 0, 0, 0, ist)
	dur := time.Since(start)
	days := int(dur.Hours()) / 24
	hours := int(dur.Hours()) % 24
	mins := int(dur.Minutes()) % 60
	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours, mins)
	}
	return fmt.Sprintf("%dh %dm", hours, mins)
}

func getKernel() string {
	b, err := os.ReadFile("/proc/version")
	if err != nil {
		return "unknown"
	}
	parts := strings.Fields(string(b))
	if len(parts) >= 3 {
		return parts[2]
	}
	return "unknown"
}

func getCPUPercent() float64 {
	read := func() (idle, total uint64) {
		f, err := os.Open("/proc/stat")
		if err != nil {
			return
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.HasPrefix(line, "cpu ") {
				continue
			}
			fields := strings.Fields(line)
			for i, v := range fields[1:] {
				n, _ := strconv.ParseUint(v, 10, 64)
				total += n
				if i == 3 {
					idle = n
				}
			}
			break
		}
		return
	}
	idle1, total1 := read()
	time.Sleep(200 * time.Millisecond)
	idle2, total2 := read()
	totalDiff := total2 - total1
	idleDiff := idle2 - idle1
	if totalDiff == 0 {
		return 0
	}
	return (1.0 - float64(idleDiff)/float64(totalDiff)) * 100.0
}

func parseMemInfo() map[string]uint64 {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil
	}
	defer f.Close()
	m := make(map[string]uint64)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) < 2 {
			continue
		}
		key := strings.TrimSuffix(parts[0], ":")
		val, _ := strconv.ParseUint(parts[1], 10, 64)
		m[key] = val
	}
	return m
}

func getMemUsed() float64 {
	m := parseMemInfo()
	if m == nil {
		return 0
	}
	return float64(m["MemTotal"]-m["MemAvailable"]) / 1024 / 1024
}

func getMemTotal() float64 {
	m := parseMemInfo()
	if m == nil {
		return 0
	}
	return float64(m["MemTotal"]) / 1024 / 1024
}