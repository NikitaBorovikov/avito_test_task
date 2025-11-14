package dto

import "avitoTestTask/internal/core/models"

// Team Request DTO
type CreateTeamRequest struct {
	TeamName string   `json:"team_name"`
	Members  []Member `json:"members"`
}

type Member struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

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

// PR Request DTO
type CreatePRRequest struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
}

type MergePRRequest struct {
	PullRequestID string `json:"pull_request_id"`
}

type ReassignPRRequest struct {
	PullRequestID string `json:"pull_request_id"`
	OldUserID     string `json:"old_reviewer_id"`
}

func (r *CreatePRRequest) ToDomainPR() models.PullRequest {
	return models.PullRequest{
		ID:       r.PullRequestID,
		Title:    r.PullRequestName,
		AuthorID: r.AuthorID,
		Status:   models.PRStatusOpen, //по умолчанию статус Open
	}
}

// User Request DTO
type SetUserActiveRequest struct {
	UserID   string `json:"user_id"`
	IsActive bool   `json:"is_active"`
}
