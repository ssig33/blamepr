package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// BlameFile runs git blame on the specified file and line number to get the commit hash.
// If lineNum is 0, it gets the commit for the most recent change to the file.
func BlameFile(filePath string, lineNum int) (string, error) {
	var cmd *exec.Cmd

	if lineNum > 0 {
		// Get the commit hash for a specific line
		cmd = exec.Command("git", "blame", "-L", fmt.Sprintf("%d,%d", lineNum, lineNum), "--porcelain", filePath)
	} else {
		// Get the commit hash for the most recent change to the file
		cmd = exec.Command("git", "log", "-n", "1", "--pretty=format:%H", filePath)
	}

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git command failed: %v", err)
	}

	commitHash, err := parseBlameOutput(output, lineNum)
	if err != nil {
		return "", err
	}

	return commitHash, nil
}

// parseBlameOutput extracts the commit hash from git blame or git log output.
func parseBlameOutput(output []byte, lineNum int) (string, error) {
	if len(output) == 0 {
		return "", fmt.Errorf("no output from git command")
	}

	outputStr := string(output)

	if lineNum > 0 {
		// Parse git blame --porcelain output
		lines := strings.Split(outputStr, "\n")
		if len(lines) == 0 {
			return "", fmt.Errorf("unexpected git blame output format")
		}

		// The first line of porcelain output has the format:
		// <commit-hash> <original-line-number> <current-line-number> <line-count>
		fields := strings.Fields(lines[0])
		if len(fields) < 1 {
			return "", fmt.Errorf("unexpected git blame output format")
		}

		return fields[0], nil
	} else {
		// For git log, the output is just the commit hash
		return strings.TrimSpace(outputStr), nil
	}
}

// GetRepoInfo returns the GitHub owner and repository name for the current directory.
func GetRepoInfo() (owner, repo string, err error) {
	// Get the remote URL
	cmd := exec.Command("git", "remote", "get-url", "origin")
	output, err := cmd.Output()
	if err != nil {
		return "", "", fmt.Errorf("failed to get remote URL: %v", err)
	}

	remoteURL := strings.TrimSpace(string(output))

	// Parse owner and repo from the URL
	if strings.HasPrefix(remoteURL, "https://github.com/") {
		// HTTPS URL format: https://github.com/owner/repo.git
		parts := strings.TrimPrefix(remoteURL, "https://github.com/")
		return parseOwnerRepo(parts)
	} else if strings.HasPrefix(remoteURL, "git@github.com:") {
		// SSH URL format: git@github.com:owner/repo.git
		parts := strings.TrimPrefix(remoteURL, "git@github.com:")
		return parseOwnerRepo(parts)
	}

	return "", "", fmt.Errorf("unsupported remote URL format: %s", remoteURL)
}

// parseOwnerRepo parses the owner and repo from a string in the format "owner/repo.git"
func parseOwnerRepo(s string) (owner, repo string, err error) {
	// Remove .git suffix if present
	s = strings.TrimSuffix(s, ".git")

	parts := strings.Split(s, "/")
	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid GitHub repository path: %s", s)
	}

	owner = parts[0]
	repo = parts[1]

	return owner, repo, nil
}