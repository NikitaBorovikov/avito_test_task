package postgres

import (
	"errors"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type PostgresRepo struct {
	UserRepo        *UserRepo
	TeamRepo        *TeamRepo
	PullRequestRepo *PullRequestRepo
}

func NewPostgresRepo(db *gorm.DB) *PostgresRepo {
	return &PostgresRepo{
		UserRepo:        NewUserRepo(db),
		TeamRepo:        NewTeamRepo(db),
		PullRequestRepo: NewPullRequestRepo(db),
	}
}

func isDublicateError(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == PostgresUniqueErrorCode
	}
	return false
}
