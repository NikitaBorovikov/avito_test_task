package postgres

import (
	"avitoTestTask/internal/core/models"
	"errors"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	PostgresUniqueErrorCode = "23505"
)

var (
	ErrDublicateTeamName = errors.New("team with that name already exists")
	ErrTeamNotFound      = errors.New("team with this ID is not found")
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
		if isDuplicateError(err) {
			return nil, ErrDublicateTeamName
		}
		return nil, err
	}
	return team, nil
}

func (r *TeamRepo) GetByName(name string) (*models.Team, error) {
	var team models.Team
	if err := r.db.Preload("Users").Where("name = ?", name).First(&team).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTeamNotFound
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
		return ErrTeamNotFound
	}
	return nil
}

func isDuplicateError(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == PostgresUniqueErrorCode
	}
	return false
}
