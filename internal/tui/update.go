package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
switch msg := msg.(type) {
	case tea.KeyMsg:
		switch { 
		case key.Matches(msg, b.keys.Quit):
			return b, tea.Quit
		case key.Matches(msg, b.keys.Enter):
      var render string
			i := b.list.SelectedItem().(SettingsConfig)
      listen := b.viewProc(i.PortApp)

      if listen {
        render = fmt.Sprintf("Listen: %s", i.PortApp)
      } else {
        render = fmt.Sprintf("Off: %s", i.PortApp)
      }

			content := lipgloss.NewStyle().Width(b.viewport.Width).Height(b.viewport.Height).Render(render)
			b.viewport.SetContent(content)
			return b, nil

		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		b.list.SetSize(msg.Width-h, msg.Height-v)

		b.viewport = viewport.New(100, 25)
	}
	var cmd tea.Cmd
	b.list, cmd = b.list.Update(msg)
	return b, cmd
}
