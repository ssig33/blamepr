package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ssig33/blamepr/pkg/git"
)

// PR represents a GitHub Pull Request.
type PR struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	URL    string `json:"html_url"`
}

// FindPRByCommit queries the GitHub API to find the PR associated with a commit.
func FindPRByCommit(commitHash string) (*PR, error) {
	// Get the GitHub token from environment or netrc
	token, err := getGitHubToken()
	if err != nil {
		return nil, err
	}

	// Get repository owner and name
	owner, repo, err := git.GetRepoInfo()
	if err != nil {
		return nil, err
	}

	// Query the GitHub API to find the PR
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits/%s/pulls", owner, repo, commitHash)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.groot-preview+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("no pull request found for commit %s", commitHash)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	var prs []PR
	if err := json.NewDecoder(resp.Body).Decode(&prs); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %v", err)
	}

	if len(prs) == 0 {
		return nil, fmt.Errorf("no pull request found for commit %s", commitHash)
	}

	return &prs[0], nil
}

// getGitHubToken returns the GitHub API token from the environment or .netrc file.
func getGitHubToken() (string, error) {
	// First, try to get the token from the environment
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		return token, nil
	}

	// If not in the environment, try to get it from .netrc
	netrcPath := filepath.Join(os.Getenv("HOME"), ".netrc")
	data, err := os.ReadFile(netrcPath)
	if err != nil {
		return "", fmt.Errorf("GitHub token not found in environment or .netrc: %v", err)
	}

	// Parse the .netrc file for the GitHub token
	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if strings.HasPrefix(line, "machine github.com") || strings.HasPrefix(line, "machine api.github.com") {
			// Look for the password/token in the next few lines
			for j := i + 1; j < len(lines) && j < i+5; j++ {
				tokenLine := strings.TrimSpace(lines[j])
				if strings.HasPrefix(tokenLine, "password") || strings.HasPrefix(tokenLine, "login") {
					parts := strings.Fields(tokenLine)
					if len(parts) >= 2 {
						return parts[1], nil
					}
				}
			}
		}
	}

	return "", fmt.Errorf("GitHub token not found in environment or .netrc")
}