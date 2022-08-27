package tui

import (
	"fmt"
	"log"
	"os/exec"
)


func viewProc(/*conf *config.Config*/) string {
  // lsof -i -P -n | grep LISTEN
  filterTCP := fmt.Sprintf("lsof -i -P -n | grep LISTEN | grep %s", "4200")
  cmd, err := exec.Command("/bin/bash", "-c", filterTCP).Output()
  if err != nil {
    log.Fatal(err)
    return "fail"
  }

  return string(cmd)
}
