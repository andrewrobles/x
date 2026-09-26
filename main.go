package main

import (
	_ "embed"
	"fmt"
	"os/exec"
	"time"
	"os"
)

func main() {
	archive()
}

func archive() error {
	markdown, err := readReminders()
	if err != nil {
		return err
	}

	filename, err := writeArchive(markdown)
	if err != nil {
		return err
	}

	fmt.Printf("Wrote %s\n", filename)

	return nil
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

func writeArchive(markdown string) (string, error) {
	filename := archiveFilename()

	err := os.WriteFile(
		filename,
		[]byte(markdown),
		0644,
	)

	if err != nil {
		return "", fmt.Errorf("write archive: %w", err)
	}

	return filename, nil
}

func archiveFilename() string {
	return time.Now().Format("2006-01-02") + ".md"
}