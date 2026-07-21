package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

const (
	totalWidth = 100
	colWidth   = 48
	gap        = 2
)

const logoText = `
  ██████╗ ███████╗██╗  ██╗ ██████╗ ██████╗  ██████╗  █████╗ ███╗   ███╗
  ██╔══██╗██╔════╝██║  ██║██╔═══██╗██╔══██╗██╔═══██╗██╔══██╗████╗ ████║
  ██████╔╝█████╗  ███████║██║   ██║██████╔╝██║   ██║███████║██╔████╔██║
  ██╔══██╗██╔══╝  ██╔══██║██║   ██║██╔══██╗██║   ██║██╔══██║██║╚██╔╝██║
  ██║  ██║███████╗██║  ██║╚██████╔╝██████╔╝╚██████╔╝██║  ██║██║ ╚═╝ ██║
  ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝`

type SysData struct {
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

type ContainerData struct {
	Name   string
	State  string
	Uptime string
}

type NetData struct {
	Interface string
	Address   string
	Tunnel    string
	Internet  string
}

type StorageEntry struct {
	Name string
	Size string
}

type StoData struct {
	Entries   []StorageEntry
	TotalUsed float64
	TotalSize float64
}

type SecData struct {
	FailedCount  int
	LastAttacker string
	LastLogin    string
}

func Render(sys SysData, containers []ContainerData, net NetData, sto StoData, sec SecData) string {
	var b strings.Builder

	b.WriteString("\n")
	b.WriteString(bold.Render(logoText) + "\n")
	b.WriteString("\n")
	b.WriteString(dim.Render("  " + strings.Repeat("─", 70)) + "\n")
	b.WriteString(dim.Render("  server.vincents.systems  ·  homeserver") + "\n")
	b.WriteString("\n\n")

	b.WriteString(lipgloss.JoinHorizontal(
		lipgloss.Top,
		renderSystem(sys),
		strings.Repeat(" ", gap),
		renderServices(containers),
	))
	b.WriteString("\n")

	b.WriteString(lipgloss.JoinHorizontal(
		lipgloss.Top,
		renderNetwork(net),
		strings.Repeat(" ", gap),
		renderSecurity(sec),
	))
	b.WriteString("\n")

	b.WriteString(lipgloss.JoinHorizontal(
		lipgloss.Top,
		renderStorage(sto),
		strings.Repeat(" ", gap),
		renderErrors(),
	))
	b.WriteString("\n\n")

	ist := time.FixedZone("IST", 5*60*60+30*60)
	now := time.Now().In(ist).Format("2006-01-02  15:04 IST")
	rule := footerRule.Render(strings.Repeat("─", 50))
	footer := footerText.Render("rehoboam  ·  v0.1.0  ·  " + now)
	b.WriteString(center(rule, totalWidth) + "\n")
	b.WriteString(center(footer, totalWidth) + "\n")
	b.WriteString("\n")

	return b.String()
}

func center(s string, width int) string {
	visible := lipgloss.Width(s)
	pad := (width - visible) / 2
	if pad < 0 {
		pad = 0
	}
	return strings.Repeat(" ", pad) + s
}

func renderSystem(s SysData) string {
	memPct := 0.0
	if s.MemTotal > 0 {
		memPct = s.MemUsed / s.MemTotal * 100
	}
	rootPct := 0.0
	if s.RootTotal > 0 {
		rootPct = s.RootUsed / s.RootTotal * 100
	}
	homePct := 0.0
	if s.HomeTotal > 0 {
		homePct = s.HomeUsed / s.HomeTotal * 100
	}
	lines := []string{
		row("Host", s.Host),
		row("Uptime", s.Uptime),
		row("Kernel", s.Kernel),
		progressRow("CPU", s.CPUPct, fmt.Sprintf("%.0f%%", s.CPUPct)),
		progressRow("Memory", memPct, fmt.Sprintf("%.1f / %.1f GB", s.MemUsed, s.MemTotal)),
		progressRow("Disk  /", rootPct, fmt.Sprintf("%.0f / %.0f GB", s.RootUsed, s.RootTotal)),
		progressRow("Disk  /home", homePct, fmt.Sprintf("%.0f / %.0f GB", s.HomeUsed, s.HomeTotal)),
	}
	return panel("SYSTEM", lines, colWidth)
}

func renderServices(cs []ContainerData) string {
	lines := make([]string, len(cs))
	for i, c := range cs {
		lines[i] = serviceRow(c.Name, c.State, c.Uptime)
	}
	return panel("SERVICES", lines, colWidth)
}

func renderNetwork(n NetData) string {
	tunnelStr := ""
	if n.Tunnel == "active" {
		tunnelStr = green.Render("●") + " cloudflared active"
	} else {
		tunnelStr = red.Render("○") + " cloudflared " + n.Tunnel
	}
	internetStr := ""
	if n.Internet == "unreachable" {
		internetStr = red.Render("○") + " unreachable"
	} else {
		internetStr = green.Render("●") + " " + n.Internet + " to 1.1.1.1"
	}
	lines := []string{
		row("Interface", n.Interface),
		row("Address", n.Address),
		row("Tunnel", tunnelStr),
		row("Internet", internetStr),
	}
	return panel("NETWORK", lines, colWidth)
}

func renderSecurity(s SecData) string {
	countStr := ""
	switch {
	case s.FailedCount == 0:
		countStr = green.Render("0") + dim.Render(" attempts today")
	case s.FailedCount < 20:
		countStr = yellow.Render(fmt.Sprintf("%d", s.FailedCount)) + dim.Render(" attempts today")
	default:
		countStr = red.Render(fmt.Sprintf("%d", s.FailedCount)) + dim.Render(" attempts today")
	}
	lines := []string{
		row("Failed SSH", countStr),
		row("Last attempt", s.LastAttacker),
		row("Last login", s.LastLogin),
	}
	return panel("SECURITY", lines, colWidth)
}

func renderStorage(s StoData) string {
	lines := make([]string, 0, len(s.Entries)+1)
	for _, e := range s.Entries {
		lines = append(lines, row(e.Name, e.Size))
	}
	lines = append(lines, row("total",
		fmt.Sprintf("%.0f GB of %.0f GB", s.TotalUsed, s.TotalSize)))
	return panel("STORAGE", lines, colWidth)
}

func renderErrors() string {
	lines := []string{dim.Render("loki integration coming soon")}
	return panel("RECENT ERRORS", lines, colWidth)
}