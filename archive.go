package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func archiveMarkdown(directory, markdown string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get home dir: %w", err)
	}

	dir := filepath.Join(home, directory)
	filename, err := writeArchive(dir, markdown)
	if err != nil {
		return "", err
	}

	if err := verifyArchive(filename, markdown); err != nil {
		return "", err
	}

	if err := commitArchive(dir); err != nil {
		return "", err
	}

	return filename, nil
}

func writeArchive(dir, markdown string) (string, error) {
	filename := archiveFilename(dir)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("create archive dir: %w", err)
	}

	file, err := os.OpenFile(
		filename,
		os.O_CREATE|os.O_WRONLY|os.O_EXCL,
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

	if err := file.Close(); err != nil {
		return "", fmt.Errorf("close archive: %w", err)
	}

	return filename, nil
}

func archiveFilename(dir string) string {
	timestamp := time.Now().Format("2006-01-02T15-04-05")
	return filepath.Join(dir, timestamp+".md")
}

func verifyArchive(filename, markdown string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read archive for verification: %w", err)
	}

	if string(data) != markdown {
		return fmt.Errorf("archive verification failed")
	}

	return nil
}

func commitArchive(dir string) error {
	if _, err := os.Stat(filepath.Join(dir, ".git")); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("inspect archive repository: %w", err)
		}
		if err := runGit(dir, "init"); err != nil {
			return err
		}
	}

	if err := runGit(dir, "add", "-A"); err != nil {
		return err
	}

	return runGit(dir, "commit", "--allow-empty-message", "-m", "")
}

func runGit(dir string, args ...string) error {
	cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git %s: %w\n%s", strings.Join(args, " "), err, output)
	}

	return nil
}
