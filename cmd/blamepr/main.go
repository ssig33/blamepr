package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ssig33/blamepr/pkg/git"
	"github.com/ssig33/blamepr/pkg/github"
	"github.com/ssig33/blamepr/pkg/browser"
)

func main() {
	openFlag := flag.Bool("open", false, "Open the PR in the default browser")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: blamepr [-open] path/to/file.go[:line]")
		os.Exit(1)
	}

	// Parse file path and optional line number
	arg := flag.Arg(0)
	filePath, lineNum, err := parseFilePathAndLine(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Get the commit hash using git blame
	commitHash, err := git.BlameFile(filePath, lineNum)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Find the PR associated with the commit
	pr, err := github.FindPRByCommit(commitHash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Output the PR information
	fmt.Printf("PR #%d: %s\n", pr.Number, pr.Title)
	fmt.Printf("URL: %s\n", pr.URL)

	// Open the PR in the browser if the -open flag is set
	if *openFlag {
		if err := browser.Open(pr.URL); err != nil {
			fmt.Fprintf(os.Stderr, "Error opening browser: %v\n", err)
			os.Exit(1)
		}
	}
}

// parseFilePathAndLine parses a string in the format "path/to/file.go:123" and returns
// the file path and line number separately.
func parseFilePathAndLine(arg string) (filePath string, lineNum int, err error) {
	parts := strings.Split(arg, ":")
	filePath = parts[0]

	if len(parts) > 1 {
		fmt.Sscanf(parts[1], "%d", &lineNum)
	}

	// Verify the file exists
	_, err = os.Stat(filePath)
	if err != nil {
		return "", 0, fmt.Errorf("file not found: %s", filePath)
	}

	return filePath, lineNum, nil
}