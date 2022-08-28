package tui

import "github.com/charmbracelet/lipgloss"

func (b Bubble) View() string {

	var view string

	view = lipgloss.JoinHorizontal(
		lipgloss.Top,
		docStyle.Render(b.list.View()),
		viewportStyle.Render(b.viewport.View()),
	)
	return view
}
