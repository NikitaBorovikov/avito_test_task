package usecases

import (
	"avitoTestTask/internal/core/models"
	"avitoTestTask/internal/core/repository"
	"fmt"
)

type TeamUC struct {
	TeamRepo repository.TeamRepo
	UserRepo repository.UserRepo
}

func NewTeamUC(teamRepo repository.TeamRepo, userRepo repository.UserRepo) *TeamUC {
	return &TeamUC{
		TeamRepo: teamRepo,
		UserRepo: userRepo,
	}
}

func (uc *TeamUC) Create(team *models.Team) (*models.Team, error) {
	createdTeam, err := uc.TeamRepo.Create(team)
	if err != nil {
		// TODO: Проверка на дубликат
		return nil, err
	}

	for i := range team.Users {
		user := &team.Users[i]
		user.TeamID = createdTeam.ID

		// Create or update user
		if err := uc.UserRepo.CreateOrUpdate(user); err != nil {
			if err := uc.TeamRepo.Delete(createdTeam.ID); err != nil {
				return nil, fmt.Errorf("failed to save user and rollback team: %v", err)
			}
			return nil, fmt.Errorf("failed to save user: %v", err)
		}
	}
	return createdTeam, nil
}

func (uc *TeamUC) GetByName(name string) (*models.Team, error) {
	return uc.TeamRepo.GetByName(name)
}
