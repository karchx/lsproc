package main

import (
	"fmt"
	"os"

	"github.com/KenethSandoval/lsproc/ui"
)

func main() {
	// call cmd.Execute()
	// cmd.ListenService()
	// TODO: mover a cmd/root
	if err := ui.NewProgram().Start(); err != nil {
		fmt.Println("Could not start ui", err)
		os.Exit(1)
	}
}
