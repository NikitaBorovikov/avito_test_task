package usecases

import (
	"avitoTestTask/internal/core/models"
	"avitoTestTask/internal/core/repository"
)

type UserUC struct {
	UserRepo repository.UserRepo
}

func NewUserUC(userRepo repository.UserRepo) *UserUC {
	return &UserUC{
		UserRepo: userRepo,
	}
}

func (uc *UserUC) CreateOrUpdate(user *models.User) error {
	return nil
}

func (uc *UserUC) GetByID(userID string) (*models.User, error) {
	return nil, nil
}

func (uc *UserUC) SetUserActive(userID string, isActive bool) (*models.User, error) {
	return nil, nil
}
