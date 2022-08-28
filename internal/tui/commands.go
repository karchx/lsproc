package tui

import (
	"fmt"
	"os/exec"
)


func (b Bubble) viewProc(portApp string) bool {
  // lsof -i -P -n | grep LISTEN
  filterTCP := fmt.Sprintf("lsof -i -P -n | grep LISTEN | grep %s", portApp)
  _, err := exec.Command("/bin/bash", "-c", filterTCP).Output()
  if err != nil {
    return false
  }
  return true
}
