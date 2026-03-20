package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// RepoInfo holds the key information about a GitHub repository.
type RepoInfo struct {
	Name        string `json:"full_name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	ForksCount  int    `json:"forks_count"`
	CreatedAt   string `json:"created_at"`
}

// parseInput extracts the owner and repository name from the input string.
// It supports two formats: "owner/repo" and "https://github.com/owner/repo".
func parseInput(input string) (string, string, error) {
	if strings.Contains(input, "github.com/") {
		parts := strings.Split(input, "github.com/")
		if len(parts) < 2 {
			return "", "", fmt.Errorf("invalid URL format")
		}
		input = parts[1]
	}

	parts := strings.Split(strings.TrimSuffix(input, "/"), "/")

	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid format, expected 'owner/repo' or GitHub URL")
	}

	return parts[0], parts[1], nil
}

// getRepo makes an HTTP GET request to the GitHub API to fetch repository details.
// It returns a RepoInfo upon success, or an error if the request fails.
func getRepo(owner string, repoName string) (RepoInfo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repoName)
	resp, err := http.Get(url)
	if err != nil {
		return RepoInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RepoInfo{}, fmt.Errorf("repository not found (status code: %d)", resp.StatusCode)
	}

	var info RepoInfo
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return RepoInfo{}, err
	}

	return info, nil
}

// formatTime parses a GitHub API date string and returns a formatted date string.
// If parsing fails, it returns the original string.
func formatTime(dateStr string) string {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return dateStr
	}

	return t.Format("15:04:05 02.01.2006")
}

func (r RepoInfo) String() string {
	formattedTime := formatTime(r.CreatedAt)

	return fmt.Sprintf(
		"Info about repository: %s\nName: %s\nDescription: %s\nStars: %d\nNumber of forks: %d\nCreatedAt: %s",
		r.Name, r.Name, r.Description, r.Stars, r.ForksCount, formattedTime,
	)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: go run cmd/github-tool/main.go -repo <owner/repo>\n")
		flag.PrintDefaults()
	}

	repoID := flag.String("repo", "", "GitHub repository name")
	flag.Parse()

	if *repoID == "" {
		fmt.Println("You did not specify the repository address.")
		return
	}

	owner, repoName, err := parseInput(*repoID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	repoInfo, err := getRepo(owner, repoName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(repoInfo)
}
