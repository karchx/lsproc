package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch {
    case key.Matches(msg, b.keys.Quit):
      return b, tea.Quit

    case key.Matches(msg, b.keys.ReloadConfig):
      var cmd tea.Cmd
      cmd = tea.EnterAltScreen
      return b, cmd
    }
  }
  return b, nil
}
