package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	project := flag.String("project", "", "Project to run (guess-number, unit-converter, qr-generator, etc.)")
	flag.Parse()

	switch *project {
	case "calculator":
		runProject("guess-number")
	default:
		fmt.Println("Invalid project")
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
