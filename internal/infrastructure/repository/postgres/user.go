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
	var users []models.User
	err := r.db.Where("team_id = ? AND is_active = ?", teamID, true).First(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) SetUserActive(userID string, isActive bool) (*models.User, error) {
	result := r.db.Model(&models.User{}).Where("id = ?", userID).Update("is_active", isActive)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, ErrUserNotFound
	}

	updatedUser, err := r.GetByID(userID)
	return updatedUser, err
}
