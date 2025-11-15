package usecases

import "avitoTestTask/internal/core/repository"

type UseCases struct {
	UserUC        *UserUC
	TeamUC        *TeamUC
	PullRequestUC *PullRequestUC
}

func NewUseCases(userRepo repository.UserRepo, teamRepo repository.TeamRepo, prRepo repository.PullRequestRepo) *UseCases {
	return &UseCases{
		UserUC:        NewUserUC(userRepo),
		TeamUC:        NewTeamUC(teamRepo, userRepo),
		PullRequestUC: NewPullRequestUC(prRepo, userRepo),
	}
}
