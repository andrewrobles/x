package main

import (
	_ "embed"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	markdown, err := readReminders()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(markdown)
}

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


func archiveFilename() string {
	return time.Now().Format("2006-01-02") + ".md"
}