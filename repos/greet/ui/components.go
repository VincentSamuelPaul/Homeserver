package ui

import (
	"fmt"
	"strings"
)

// sectionTitle renders a panel title line e.g. "─ SYSTEM ─────"
func sectionTitle(title string, width int) string {
	label := " " + title + " "
	dashes := width - len(label) - 2
	if dashes < 0 {
		dashes = 0
	}
	return dim.Render("─") + bright.Render(label) + dim.Render(strings.Repeat("─", dashes))
}

// row renders a label + value line inside a panel
func row(label, value string) string {
	l := dim.Render(fmt.Sprintf("%-16s", label))
	v := bright.Render(value)
	return l + v
}

// statusDot returns colored dot + symbol for a container state
func statusDot(state string) string {
	switch state {
	case "running":
		return green.Render("●")
	case "stopped", "exited":
		return red.Render("○")
	case "restarting":
		return yellow.Render("~")
	default:
		return dim.Render("?")
	}
}

// serviceRow renders one container line
func serviceRow(name, state, uptime string) string {
	dot := statusDot(state)
	n := bright.Render(fmt.Sprintf("%-18s", name))
	st := dim.Render(fmt.Sprintf("%-12s", state))
	u := dim.Render(uptime)
	return n + dot + "  " + st + u
}

// progressBar renders an ASCII progress bar (no unicode blocks)
func progressBar(pct float64, width int) string {
	filled := int(pct / 100.0 * float64(width))
	if filled > width {
		filled = width
	}
	empty := width - filled
	bar := barFill.Render(strings.Repeat("#", filled)) +
		barEmpty.Render(strings.Repeat("-", empty))
	return "[" + bar + "]"
}

// progressRow renders a label + bar + value line
func progressRow(label string, pct float64, valueStr string) string {
	l := dim.Render(fmt.Sprintf("%-16s", label))
	bar := progressBar(pct, 12)
	v := bright.Render("  " + valueStr)
	return l + bar + v
}

// panel wraps lines in a rounded border box with a title
func panel(title string, lines []string, width int) string {
	inner := width - 4 // account for border + padding

	var content strings.Builder
	content.WriteString(sectionTitle(title, inner) + "\n")
	content.WriteString("\n")
	for _, l := range lines {
		content.WriteString(l + "\n")
	}

	return panelStyle.Width(width).Render(content.String())
}

// errorRow renders a service name + message line
func errorRow(service, msg string) string {
	svc := red.Render(fmt.Sprintf("%-14s", service))
	m := dim.Render(msg)
	return svc + m
}