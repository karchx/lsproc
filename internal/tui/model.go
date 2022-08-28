package tui

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/karchx/lsproc/internal/config"
)

type sessionState int

type SettingsConfig struct {
	NameApp        string
	PortApp        string
	Command        string
	PathApp        string
	PrivilegesRoot bool
}

type Bubble struct {
  state sessionState
  config config.Config
  keys KeyMap
  list list.Model
  items []list.Item
  viewport viewport.Model
  listen bool
}

func (i SettingsConfig) Title() string       { return i.NameApp }
func (i SettingsConfig) Description() string { return i.Command }
func (i SettingsConfig) FilterValue() string { return i.NameApp }

func New() Bubble {
  cfg, err := config.ParseConfig()

  if err != nil {
    log.Fatal(err)
  }

var procs []list.Item
	for _, s := range cfg.Services.Containers {
		ap := SettingsConfig{NameApp: s.NameApp, Command: s.Command, PathApp: s.PathApp, PortApp: s.PortApp}
		procs = append(procs, ap)
	}

	listModel := list.New(procs, list.NewDefaultDelegate(), 0, 0)
	listModel.Title = "Services list"

  return Bubble{
    config: cfg,
    keys: DefaultKeyMap(),
    list: listModel,
    listen: true,
  }
}
