package main

import (
	"fmt"
	"os"
)

var version = "0.0.0"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`Usage:
  x <command>
  x -v | --version

Commands:
  reminders  Archive reminders, commit, then delete them
  notepad    Archive ~/Notepad/.md, commit, then empty it

Options:
  -v, --version  Print version`)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-v", "--version":
		fmt.Println(version)
	case "reminders":
		if err := archiveReminders(); err != nil {
			fmt.Println("reminders:", err)
			os.Exit(1)
		}
	case "notepad":
		if err := archiveNotepad(); err != nil {
			fmt.Println("notepad:", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
