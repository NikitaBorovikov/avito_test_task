package models

type User struct {
	ID          string        `gorm:"primaryKey;type:varchar(100)"`
	Name        string        `gorm:"type:varchar(255);not null"`
	TeamID      uint          `gorm:"not null;index"`
	IsActive    bool          `gorm:"default:true"`
	Team        *Team         `gorm:"foreignKey:TeamID"`
	AuthoredPRs []PullRequest `gorm:"foreignKey:AuthorID"`
}
