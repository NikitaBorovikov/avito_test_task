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
	author, err := uc.UserRepo.GetByID(pr.AuthorID)
	if err != nil {
		return nil, err
	}

	reviewers, err := uc.setReviewers(author.TeamID, author.ID)
	if err != nil {
		return nil, err
	}

	pr.Reviewers = reviewers
	updatedPR, err := uc.PullRequestRepo.Create(pr)
	if err != nil {
		return nil, err
	}
	return updatedPR, nil
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

func (uc *PullRequestUC) findRandomReplacement(authorID string, teamID uint, existingRev []string) (string, error) {
	activeUsers, err := uc.UserRepo.GetActiveUsersByTeam(teamID)
	if err != nil {
		return "", err
	}

	candidates := make([]string, 0, len(activeUsers))
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

func (uc *PullRequestUC) setReviewers(teamID uint, authorID string) ([]string, error) {
	activeUsers, err := uc.UserRepo.GetActiveUsersByTeam(teamID)
	if err != nil {
		return nil, err
	}

	candidates := make([]string, 0, len(activeUsers))
	for _, user := range activeUsers {
		if user.ID != authorID {
			candidates = append(candidates, user.ID)
		}
	}

	if len(candidates) == 0 {
		return nil, ErrNoCandidate
	}

	rand.Shuffle(len(candidates), func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	maxReviewers := min(2, len(candidates))
	return candidates[:maxReviewers], nil
}

func isUserInReviewers(userID string, reviewers []string) bool {
	for _, id := range reviewers {
		if id == userID {
			return true
		}
	}
	return false
}
