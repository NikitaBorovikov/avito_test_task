package postgres

import (
	"avitoTestTask/internal/core/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// Если пользователь уже сущесвует, то обновляем его данные. Если не сущесвует - создаем
func (r *UserRepo) CreateOrUpdate(user *models.User) error {
	return r.db.Save(user).Error
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
