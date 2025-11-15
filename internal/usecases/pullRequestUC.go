package usecases

import (
	"avitoTestTask/internal/core/models"
	"avitoTestTask/internal/core/repository"
	"errors"
	"math/rand"
)

var (
	ErrAlreadyMerged = errors.New("already merged")
	ErrNoCandidate   = errors.New("there aren't available candidates")
	ErrNotAssigned   = errors.New("user is not assigned to this pull request")
)

type PullRequestUC struct {
	PullRequestRepo repository.PullRequestRepo
	UserRepo        repository.UserRepo
}

func NewPullRequestUC(pullRequestRepo repository.PullRequestRepo, userRepo repository.UserRepo) *PullRequestUC {
	return &PullRequestUC{
		PullRequestRepo: pullRequestRepo,
		UserRepo:        userRepo,
	}
}

func (uc *PullRequestUC) Create(pr *models.PullRequest) (*models.PullRequest, error) {
	// При создании PR автоматически назначаются до двух активных ревьюверов из команды автора, исключая самого автора.
	return uc.PullRequestRepo.Create(pr)
}

func (uc *PullRequestUC) GetByReviewer(userID string) ([]models.PullRequest, error) {
	return uc.PullRequestRepo.GetByReviewer(userID)
}

func (uc *PullRequestUC) Merge(prID string) (*models.PullRequest, error) {
	return uc.PullRequestRepo.Merge(prID)
}

func (uc *PullRequestUC) Reassign(prID, oldUserID string) (*models.PullRequest, error) {
	pr, err := uc.PullRequestRepo.GetByID(prID)
	if err != nil {
		return nil, err
	}
	if pr.Status == models.PRStatusMerged {
		return nil, ErrAlreadyMerged
	}
	if !isUserInReviewers(oldUserID, pr.Reviewers) {
		return nil, ErrNotAssigned
	}

	oldUser, err := uc.UserRepo.GetByID(oldUserID)
	if err != nil {
		return nil, err
	}

	newReviewerID, err := uc.findRandomReplacement(pr.AuthorID, oldUser.TeamID, pr.Reviewers)
	if err != nil {
		return nil, err
	}

	updatedPR, err := uc.PullRequestRepo.Reassign(prID, oldUser.ID, newReviewerID)
	if err != nil {
		return nil, err
	}
	return updatedPR, nil
}

func isUserInReviewers(userID string, reviewers []string) bool {
	for _, id := range reviewers {
		if id == userID {
			return true
		}
	}
	return false
}

func (uc *PullRequestUC) findRandomReplacement(authorID string, teamID uint, existingRev []string) (string, error) {
	activeUsers, err := uc.UserRepo.GetActiveUsersByTeam(teamID)
	if err != nil {
		return "", err
	}

	candidates := make([]string, 0, 2)
	for _, user := range activeUsers {
		if user.ID != authorID && !isUserInReviewers(user.ID, existingRev) {
			candidates = append(candidates, user.ID)
		}
	}

	if len(candidates) == 0 {
		return "", ErrNoCandidate
	}
	return candidates[rand.Intn(len(candidates))], nil
}
