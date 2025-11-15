package repository

import "avitoTestTask/internal/core/models"

type UserRepo interface {
	CreateOrUpdate(user *models.User) error
	GetByID(userID string) (*models.User, error)
	SetUserActive(userID string, isActive bool) (*models.User, error)
	GetActiveUsersByTeam(teamID uint) ([]models.User, error)
}
