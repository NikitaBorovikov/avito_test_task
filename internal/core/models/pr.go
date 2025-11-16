package models

import "time"

type PRStatus string

const (
	PRStatusMerged PRStatus = "MERGED"
	PRStatusOpen   PRStatus = "OPEN"
)

type PullRequest struct {
	ID        string     `gorm:"primaryKey;type:varchar(100)"`
	Name      string     `gorm:"type:varchar(500);not null"`
	Status    PRStatus   `gorm:"type:varchar(64);not null;default:'OPEN'"`
	AuthorID  string     `gorm:"type:varchar(100);not null;index"`
	Reviewers []string   `gorm:"type:jsonb;serializer:json"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	MergedAt  *time.Time `gorm:"null"`
	Author    *User      `gorm:"foreignKey:AuthorID"`
}
