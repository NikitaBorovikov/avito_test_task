package dto

import (
	"avitoTestTask/internal/core/models"
	"time"
)

// Main Response DTO
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code, msg string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: msg,
	}
}

// Team Response DTO
type CreateTeamOKResponse struct {
	Team Team `json:"team"`
}

type GetTeamByNameResponse struct {
	TeamName string   `json:"team_name"`
	Members  []Member `json:"members"`
}

type Team struct {
	TeamName string   `json:"team_name"`
	Members  []Member `json:"members"`
}

func NewCreateTeamOKResponse(team *models.Team) *CreateTeamOKResponse {
	if team == nil {
		return nil
	}
	members := convertUsersToMembers(team.Users)
	return &CreateTeamOKResponse{
		Team: Team{
			TeamName: team.Name,
			Members:  members,
		},
	}
}

func NewGetTeamByNameResponse(team *models.Team) *GetTeamByNameResponse {
	if team == nil {
		return nil
	}
	members := convertUsersToMembers(team.Users)
	return &GetTeamByNameResponse{
		TeamName: team.Name,
		Members:  members,
	}
}

func convertUsersToMembers(users []models.User) []Member {
	if len(users) == 0 {
		return []Member{}
	}

	members := make([]Member, 0, len(users))
	for _, user := range users {
		member := Member{
			UserID:   user.ID,
			Username: user.Name,
			IsActive: user.IsActive,
		}
		members = append(members, member)
	}
	return members
}

// PR Response DTO
type PullRequestShort struct {
	PullRequestID   string          `json:"pull_request_id"`
	PullRequestName string          `json:"pull_request_name"`
	AuthorID        string          `json:"author_id"`
	Status          models.PRStatus `json:"status"`
}

type CreatePRResponse struct {
	PullRequest PullRequest `json:"pr"`
}

type ReassignPRResponse struct {
	PullRequest PullRequest `json:"pr"`
}

type MergePRResponse struct {
	PullRequest PullRequest `json:"pr"`
}

type PullRequest struct {
	PullRequestID     string          `json:"pull_request_id"`
	PullRequestName   string          `json:"pull_request_name"`
	AuthorID          string          `json:"author_id"`
	Status            models.PRStatus `json:"status"`
	AssignedReviewers []string        `json:"assigned_reviewers"`
	MergedAt          time.Time       `json:"mergedAt,omitempty"`
}

func NewCreatePRResponse(pr *models.PullRequest) *CreatePRResponse {
	if pr == nil {
		return nil
	}
	return &CreatePRResponse{
		PullRequest: PullRequest{
			PullRequestID:     pr.ID,
			PullRequestName:   pr.Title,
			AuthorID:          pr.AuthorID,
			Status:            pr.Status,
			AssignedReviewers: pr.Reviewers,
		},
	}
}

func NewMergePRResponse(pr *models.PullRequest) *MergePRResponse {
	if pr == nil {
		return nil
	}
	return &MergePRResponse{
		PullRequest: PullRequest{
			PullRequestID:     pr.ID,
			PullRequestName:   pr.Title,
			AuthorID:          pr.AuthorID,
			Status:            pr.Status,
			AssignedReviewers: pr.Reviewers,
			MergedAt:          pr.MergedAt,
		},
	}
}

func NewReassignPRResponse(pr *models.PullRequest) *ReassignPRResponse {
	if pr == nil {
		return nil
	}
	return &ReassignPRResponse{
		PullRequest: PullRequest{
			PullRequestID:     pr.ID,
			PullRequestName:   pr.Title,
			AuthorID:          pr.AuthorID,
			Status:            pr.Status,
			AssignedReviewers: pr.Reviewers,
		},
	}
}

// User Response DTO
type GetUserReviewPRResponse struct {
	UserID       string             `json:"user_id"`
	PullRequests []PullRequestShort `json:"pull_requests"`
}

type SetUserActiveResponse struct {
	User User `json:"user"`
}

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"is_active"`
}

func NewGetUserReviewPRResponse(userID string, prs []models.PullRequest) *GetUserReviewPRResponse {
	pullRequests := make([]PullRequestShort, 0, len(prs))

	for _, p := range prs {
		pullRequest := PullRequestShort{
			PullRequestID:   p.ID,
			PullRequestName: p.Title,
			AuthorID:        p.AuthorID,
			Status:          p.Status,
		}
		pullRequests = append(pullRequests, pullRequest)
	}
	return &GetUserReviewPRResponse{
		UserID:       userID,
		PullRequests: pullRequests,
	}
}

func NewSetUserActiveResponse(teamName string, user *models.User) *SetUserActiveResponse {
	if user == nil {
		return nil
	}
	return &SetUserActiveResponse{
		User: User{
			UserID:   user.ID,
			Username: user.Name,
			TeamName: teamName,
			IsActive: user.IsActive,
		},
	}
}
