package main

import (
	"log"
	"os/exec"
)

func diff() string {
	output := executeCommand()
	return string(output)
}

func executeCommand() []byte {
	output, err := exec.Command("git", "diff", "--color").Output();
	if err != nil {
		log.Panicln(err)
	}

	return output
}
