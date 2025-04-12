package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

var mainColor = color.New(color.FgHiCyan)
var cmdColor = color.New(color.FgHiMagenta)
var commentColor = color.New(color.FgHiGreen)

// Get the first comment line from a file
func getFirstComment(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		return ""
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//") {
			return strings.TrimSpace(strings.TrimPrefix(line, "//"))
		}
	}
	return ""
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory_path>")
		fmt.Println("Example: go run main.go types/slices")
		os.Exit(1)
	}

	dirPath := os.Args[1]

	// Find all Go files that start with "example"
	files, err := filepath.Glob(filepath.Join(dirPath, "example*.go"))
	if err != nil {
		fmt.Printf("Error searching for example files: %v\n", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Printf("No example files found in %s\n", dirPath)
		os.Exit(1)
	}

	// Display available examples
	mainColor.Println("\nAvailable examples:")
	for i, file := range files {
		filename := filepath.Base(file)
		comment := getFirstComment(file)
		mainColor.Printf("%d. %s", i+1, filename)
		if comment != "" {
			commentColor.Printf(" %s", comment)
		}
		fmt.Println()
	}

	// Get user selection
	var choice int
	mainColor.Print("\nSelect an example to run (1-", len(files), "): ")
	fmt.Scanf("%d", &choice)

	if choice < 1 || choice > len(files) {
		fmt.Println("Invalid choice")
		os.Exit(1)
	}

	selectedFile := files[choice-1]
	mainColor.Printf("\nRunning %s...\n\n", filepath.Base(selectedFile))

	// Run the example in its original directory
	cmd := fmt.Sprintf("go run -C %s %s", dirPath, filepath.Base(selectedFile))
	mainColor.Print("Command: ")
	cmdColor.Println(cmd)
	mainColor.Println("Output:")
	mainColor.Println("----------------------------------------")

	// Execute the command
	execCmd := exec.Command("go", "run", filepath.Base(selectedFile))
	execCmd.Dir = dirPath
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	if err := execCmd.Run(); err != nil {
		fmt.Printf("Error running example: %v\n", err)
		os.Exit(1)
	}

	mainColor.Println("----------------------------------------")
	mainColor.Println("Example completed")
}
