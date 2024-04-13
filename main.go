package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Welcome to go mini-projects")
	fmt.Println("Please type guess-number, unit-converter or qr-generator to access a project")

	for {
		var project string
		fmt.Scanln(&project)

		if project != "guess-number" && project != "unit-converter" && project != "qr-generator" {
			fmt.Println("Please enter a valid project")
			continue
		} else {
			runProject(project)
			break
		}
	}
}

func runProject(projectName string) {
	cmd := exec.Command("go", "run", projectName+"/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running project:", err)
	}
}
