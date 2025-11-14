package dto

import "avitoTestTask/internal/core/models"

type CreateTeamRequest struct {
	TeamName string   `json:"team_name"`
	Members  []Member `json:"members"`
}

type Member struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

// TODO: нужно сделать валидацию required полей

func (r *CreateTeamRequest) ToDomainTeam() models.Team {
	users := make([]models.User, 0, len(r.Members))

	for _, member := range r.Members {
		user := models.User{
			ID:       member.UserID,
			Name:     member.Username,
			IsActive: member.IsActive,
		}
		users = append(users, user)
	}

	return models.Team{
		Name:  r.TeamName,
		Users: users,
	}
}
