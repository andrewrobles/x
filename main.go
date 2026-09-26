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

	file, err := os.OpenFile(
		filename,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return "", fmt.Errorf("open archive: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(markdown); err != nil {
		return "", fmt.Errorf("write archive: %w", err)
	}

	if err := file.Sync(); err != nil {
		return "", fmt.Errorf("sync archive: %w", err)
	}

	return filename, nil
}

func archiveFilename() string {
	return time.Now().Format("2006-01-02") + ".md"
}