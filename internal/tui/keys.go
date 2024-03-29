package tui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
  Quit         key.Binding
	Exit         key.Binding
	ToggleBox    key.Binding
   Enter key.Binding
}

func DefaultKeyMap() KeyMap {
  return KeyMap{
    Quit: key.NewBinding(
      key.WithKeys("ctrl+c"),
    ),
    Exit: key.NewBinding(
      key.WithKeys("q"),
    ),
    ToggleBox: key.NewBinding(
      key.WithKeys("tab"),
    ),
    Enter: key.NewBinding(
      key.WithKeys("enter"),
    ),
  }
}
