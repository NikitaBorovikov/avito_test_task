package postgres

import "avitoTestTask/internal/core/models"

type TeamRepo struct {
}

func NewTeamRepo() *TeamRepo {
	return &TeamRepo{}
}

func (r *TeamRepo) Create(team *models.Team) (*models.Team, error) {
	return nil, nil
}

func (r *TeamRepo) GetByName(name string) (*models.Team, error) {
	return nil, nil
}

func (r *TeamRepo) Delete(teamID uint) error {
	return nil
}
