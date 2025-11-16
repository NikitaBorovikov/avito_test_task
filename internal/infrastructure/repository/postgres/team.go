package postgres

import (
	apperrors "avitoTestTask/internal/appErrors"
	"avitoTestTask/internal/core/models"
	"errors"

	"gorm.io/gorm"
)

type TeamRepo struct {
	db *gorm.DB
}

func NewTeamRepo(db *gorm.DB) *TeamRepo {
	return &TeamRepo{
		db: db,
	}
}

func (r *TeamRepo) Create(team *models.Team) (*models.Team, error) {
	if err := r.db.Create(team).Error; err != nil {
		if isDublicateError(err) {
			return nil, apperrors.ErrDuplicateTeamName
		}
		return nil, err
	}
	return team, nil
}

func (r *TeamRepo) GetByName(name string) (*models.Team, error) {
	var team models.Team
	if err := r.db.Preload("Users").Where("name = ?", name).First(&team).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrTeamNotFound
		}
		return nil, err
	}
	return &team, nil
}

func (r *TeamRepo) GetByID(teamID uint) (*models.Team, error) {
	var team models.Team
	if err := r.db.Where("id = ?", teamID).First(&team).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, err
	}
	return &team, nil
}

func (r *TeamRepo) Delete(teamID uint) error {
	result := r.db.Where("id = ?", teamID).Delete(&models.Team{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrTeamNotFound
	}
	return nil
}
