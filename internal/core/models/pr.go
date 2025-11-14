package models

import "time"

type PRStatus string

const (
	PRStatusMerged PRStatus = "MERGED"
	PRStatusOpen   PRStatus = "OPEN"
)

type PullRequest struct {
	ID        string
	Title     string
	Status    PRStatus
	AuthorID  string
	Reviewers []string
	CreatedAt time.Time
}
