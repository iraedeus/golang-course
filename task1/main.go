package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
)

// RepoInfo holds the key information about a GitHub repository.
// It includes JSON tags for automatic unmarshaling from the GitHub API.
type RepoInfo struct {
	Name        string `json:"full_name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	ForksCount  int    `json:"forks_count"`
	CreatedAt   string `json:"created_at"`
}

// parseInput extracts the owner and repository name from the input string.
// It supports both "owner/repo" format and full GitHub URLs.
func parseInput(input string) (string, string) {
	input = strings.TrimSuffix(input, "/")
	parts := strings.Split(input, "/")

	if len(parts) < 2 {
		return "", ""
	}

	repo := parts[len(parts)-1]
	owner := parts[len(parts)-2]

	return owner, repo
}

// getRepo makes an HTTP GET request to the GitHub API to fetch repository details.
// It returns a pointer to RepoInfo upon success, or an error if the request fails.
func getRepo(owner string, repoName string) (*RepoInfo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repoName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("repository not found (status code: %d)", resp.StatusCode)
	}

	var info RepoInfo
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// printRepoInfo formats and prints the repository details to the standard output.
func printRepoInfo(repoInfo *RepoInfo) {
	fmt.Printf("Name: %s\n", repoInfo.Name)
	fmt.Printf("Description: %s\n", repoInfo.Description)
	fmt.Printf("Stars: %d\n", repoInfo.Stars)
	fmt.Printf("Number of forks: %d\n", repoInfo.ForksCount)
	fmt.Printf("CreatedAt: %s\n", repoInfo.CreatedAt)
}

func main() {
	repoID := flag.String("repo", "", "GitHub repository name")
	flag.Parse()

	if *repoID == "" {
		fmt.Println("You did not specify the repository address.")
		return
	}

	owner, repoName := parseInput(*repoID)
	repoInfo, err := getRepo(owner, repoName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Info about repository: %s\n\n", *repoID)
	printRepoInfo(repoInfo)
}
