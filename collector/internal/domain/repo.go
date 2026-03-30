package domain

import "errors"

type Repo struct {
	Name        string // Full name of the repository (owner/repo)
	Description string // Short description of the repository
	Stars       int    // Number of stargazers
	Forks       int    // Number of forks
	CreatedAt   string // Creation date in RFC3339 format
}

var ErrNoFound = errors.New("repository no found")
