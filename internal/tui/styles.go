package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var (
	color    = termenv.EnvColorProfile().Color
	keyword  = termenv.Style{}.Foreground(color("204")).Background(color("235")).Styled
	help     = termenv.Style{}.Foreground(color("241")).Styled
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	viewportStyle = lipgloss.NewStyle().
			Margin(0, 0, 0, 0).
			Padding(1, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)
)
