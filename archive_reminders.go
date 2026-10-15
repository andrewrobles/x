package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func archiveReminders() error {
	markdown, err := readReminders()
	if err != nil {
		return err
	}

	if strings.TrimSpace(markdown) == "" {
		fmt.Println("No reminders to archive.")
		return nil
	}

	filename, err := archiveMarkdown("Reminders", markdown)
	if err != nil {
		return err
	}

	if err := deleteReminders(); err != nil {
		return err
	}

	fmt.Printf("Archived reminders to %s\n", filename)
	return nil
}

//go:embed scripts/read-reminders/read-reminders.applescript
var readRemindersScript string

type reminderList struct {
	Name      string         `json:"name"`
	Reminders []reminderItem `json:"reminders"`
}

type reminderItem struct {
	Name      string `json:"name"`
	Notes     string `json:"notes"`
	Completed bool   `json:"completed"`
}

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

	var lists []reminderList
	if err := json.Unmarshal(output, &lists); err != nil {
		return "", fmt.Errorf("parse reminders: %w\n%s", err, output)
	}

	return remindersMarkdown(lists), nil
}

func remindersMarkdown(lists []reminderList) string {
	var markdown strings.Builder

	for _, list := range lists {
		markdown.WriteString(list.Name + "  \n")

		for _, item := range list.Reminders {
			checkbox := "- [ ] "
			if item.Completed {
				checkbox = "- [x] "
			}
			markdown.WriteString(checkbox + item.Name + "  \n")

			if notes := strings.TrimSpace(item.Notes); notes != "" {
				for _, line := range strings.Split(normalizeNewlines(item.Notes), "\n") {
					markdown.WriteString("  > " + line + "  \n")
				}
			}
		}

		markdown.WriteString("\n")
	}

	return markdown.String()
}

func normalizeNewlines(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	return s
}

//go:embed scripts/delete-reminders.applescript
var deleteRemindersScript string

func deleteReminders() error {
	cmd := exec.Command("osascript", "-e", deleteRemindersScript)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"delete reminders: %w\n%s",
			err,
			output,
		)
	}

	return nil
}
