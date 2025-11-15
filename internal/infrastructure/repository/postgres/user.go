package postgres

import "avitoTestTask/internal/core/models"

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) CreateOrUpdate(user *models.User) error {
	return nil
}

func (r *UserRepo) GetByID(userID string) (*models.User, error) {
	return nil, nil
}

func (r *UserRepo) GetActiveUsersByTeam(teamID uint) ([]models.User, error) {
	return nil, nil
}

func (uc *UserRepo) SetUserActive(userID string, isActive bool) (*models.User, error) {
	return nil, nil
}
