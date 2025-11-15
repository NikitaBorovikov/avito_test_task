package repository

import "avitoTestTask/internal/core/models"

type TeamRepo interface {
	Create(team *models.Team) (*models.Team, error)
	GetByName(name string) (*models.Team, error)
	Delete(teamID uint) error
}
