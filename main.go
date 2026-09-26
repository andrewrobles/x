package main

import (
	_ "embed"
	"fmt"
	"os/exec"
)

//go:embed scripts/read-reminders.applescript
var readRemindersScript string

func readReminders() (string, error) {
	cmd := exec.Command("osascript", "-e", readRemindersScript)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf(
			"read reminders: %w\n%s",
			err,
			output,
		)
	}

	return string(output), nil
}

func main() {
	markdown, err := readReminders()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(markdown)
}