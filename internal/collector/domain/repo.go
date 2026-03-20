package domain

// Repo represents the core domain model of a GitHub repository.
// It is independent of any external APIs or transport protocols.
type Repo struct {
	Name        string // Full name of the repository (owner/repo)
	Description string // Short description of the repository
	Stars       int    // Number of stargazers
	Forks       int    // Number of forks
	CreatedAt   string // Creation date in RFC3339 format
}
