package postgres

import (
	"avitoTestTask/internal/core/models"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user with this ID is not found")
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
	var user models.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetActiveUsersByTeam(teamID uint) ([]models.User, error) {
	return nil, nil
}

func (uc *UserRepo) SetUserActive(userID string, isActive bool) (*models.User, error) {
	return nil, nil
}
