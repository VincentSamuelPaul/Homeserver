package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Text hierarchy
	bright = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	dim    = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888"))
	bold   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true)

	// Status colors
	green  = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	red    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00"))

	// Panel borders — white border, padded inside
	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFFFFF")).
			Padding(0, 1)

	// Footer
	footerRule = lipgloss.NewStyle().Foreground(lipgloss.Color("#444444"))
	footerText = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))

	// Progress bar characters — using ASCII to avoid font issues
	barFill  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	barEmpty = lipgloss.NewStyle().Foreground(lipgloss.Color("#444444"))
)