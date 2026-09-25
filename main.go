package main

import (

	_ "embed"
	"fmt"
	"os/exec"

)

//go:embed scripts/archive.applescript

var archiveScript string

func archive() error {

	cmd := exec.Command("osascript", "-e", archiveScript)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("run archive script: %w", err)
	}
	fmt.Print(string(output))
	return nil

}

func main() {
	if err := archive(); err != nil {
		fmt.Println(err)
	}
}