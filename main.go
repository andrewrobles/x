package main

import (
	_ "embed"
	"fmt"
	"os/exec"
	"time"
	"os"
	"strings"
	"path/filepath"
)

var version = "0.0.0"

func main() {
    if len(os.Args) < 2 {
        fmt.Println("usage: x <command>")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "-v", "--version":
        fmt.Println(version)
    case "archive":
        if err := archive(); err != nil {
            fmt.Println("archive:", err)
            os.Exit(1)
        }
    default:
        fmt.Printf("unknown command: %s\n", os.Args[1])
        os.Exit(1)
    }
}

func archive() error {
	// 1. Read
	markdown, err := readReminders()
	if err != nil {
		return err
	}

	if strings.TrimSpace(markdown) == "" {
		fmt.Println("No reminders to archive.")
		return nil
	}

	// 2. Write
	filename, err := writeArchive(markdown)
	if err != nil {
		return err
	}

	// 3. Verify
	if err := verifyArchive(filename, markdown); err != nil {
		return err
	}

	// 4. Commit
    if err := commitArchive(); err != nil {
        return err
    }

	// 5. Delete
	if err := deleteReminders(); err != nil {
		return err
	}

	fmt.Printf("Archived reminders to %s\n", filename)

	return nil
}

func commitArchive() error {
    home, err := os.UserHomeDir()
    if err != nil {
        return fmt.Errorf("get home dir: %w", err)
    }
    dir := filepath.Join(home, "Reminders")

    if _, err := os.Stat(filepath.Join(dir, ".git")); os.IsNotExist(err) {
        if err := runGit(dir, "init"); err != nil {
            return err
        }
    }

    if err := runGit(dir, "add", "-A"); err != nil {
        return err
    }

    // git rejects "" as a message unless explicitly allowed
    if err := runGit(dir, "commit", "--allow-empty-message", "-m", ""); err != nil {
        return err
    }

    return nil
}

func runGit(dir string, args ...string) error {
    cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)

    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("git %s: %w\n%s", strings.Join(args, " "), err, output)
    }

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
    filename, err := archiveFilename()
    if err != nil {
        return "", err
    }

    if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
        return "", fmt.Errorf("create archive dir: %w", err)
    }

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

func archiveFilename() (string, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return "", fmt.Errorf("get home dir: %w", err)
    }
    return filepath.Join(home, "Reminders", time.Now().Format("2006-01-02")+".md"), nil
}

func verifyArchive(filename, markdown string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read archive for verification: %w", err)
	}

	if !strings.Contains(string(data), markdown) {
		return fmt.Errorf("archive verification failed")
	}

	return nil
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