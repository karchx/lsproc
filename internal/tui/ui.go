package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/karchx/lsproc/internal/client"
	"github.com/karchx/lsproc/internal/config"
	"github.com/muesli/termenv"
)

type item struct {
	title, desc string
}

type SettingsConfig struct {
	NameApp        string
	PortApp        string
	Command        string
	PathApp        string
	PrivilegesRoot bool
}

func (i SettingsConfig) Title() string       { return i.NameApp }
func (i SettingsConfig) Description() string { return i.Command }
func (i SettingsConfig) FilterValue() string { return i.NameApp }

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

type model struct {
	list     list.Model
	items    []list.Item
	viewport viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			i := m.list.SelectedItem().(SettingsConfig)
			out, err := client.RunCommand(i.Command, i.PathApp)
			if err != nil {
				fmt.Printf("ERROR: %v ", err.Error())
			}
			content := lipgloss.NewStyle().Width(m.viewport.Width).Height(m.viewport.Height).Render(out)
			m.viewport.SetContent(content)
			return m, nil

		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

		m.viewport = viewport.New(100, 25)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var view string

	view = lipgloss.JoinHorizontal(
		lipgloss.Top,
		docStyle.Render(m.list.View()),
		viewportStyle.Render(m.viewport.View()),
	)
	return view
}

func NewProgram() *tea.Program {
	cfg, err := config.ParseConfig()

	if err != nil {
		fmt.Print(err)
	}

	var procs []list.Item
	for _, s := range cfg.Services.Containers {
		ap := SettingsConfig{NameApp: s.NameApp, Command: s.Command, PathApp: s.PathApp}
		procs = append(procs, ap)
	}

	m := model{list: list.New(procs, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Services list"
	return tea.NewProgram(m, tea.WithAltScreen())
}
