package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gdbinit/mini"
)

func die(err error) {
	os.Stdout.WriteString("\x1b[2J") // clear the screen
	os.Stdout.WriteString("\x1b[H")  // reposition the cursor
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func main() {
	var editor mini.Editor

	if err := editor.Init(); err != nil {
		die(err)
	}
	defer editor.Close()

	if len(os.Args) > 1 {
		err := editor.OpenFile(os.Args[1])
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			die(err)
		}
	}

	editor.SetStatusMessage("HELP: Ctrl-S = save | Ctrl-Q = quit | Ctrl-F = find")

	for {
		editor.Render()
		if err := editor.ProcessKey(); err != nil {
			if err == mini.ErrQuitEditor {
				break
			}
			die(err)
		}
	}
}
