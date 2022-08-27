package tui

import (
	"fmt"
	"log"

	"github.com/karchx/lsproc/internal/config"
)

type sessionState int

type Bubble struct {
  state sessionState
  config config.Config
  keys KeyMap
}

func New() Bubble {
  cfg, err := config.ParseConfig()
  fmt.Println(cfg.Services.Containers)
  if err != nil {
    log.Fatal(err)
  }
  return Bubble{
    config: cfg,
    keys: DefaultKeyMap(),
  }
}
