package postgres

import (
	apperrors "avitoTestTask/internal/appErrors"
	"avitoTestTask/internal/core/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PullRequestRepo struct {
	db *gorm.DB
}

func NewPullRequestRepo(db *gorm.DB) *PullRequestRepo {
	return &PullRequestRepo{
		db: db,
	}
}

func (r *PullRequestRepo) Create(pr *models.PullRequest) (*models.PullRequest, error) {
	if err := r.db.Create(pr).Error; err != nil {
		if isDuplicateError(err) {
			return nil, apperrors.ErrDuplicatePRID
		}
		return nil, err
	}
	return pr, nil
}

func (r *PullRequestRepo) GetByReviewer(userID string) ([]models.PullRequest, error) {
	var prs []models.PullRequest
	err := r.db.
		Preload("Author").
		Preload("Author.Team").
		Where("reviewers @> ?", fmt.Sprintf(`["%s"]`, userID)).
		Find(&prs).Error

	if err != nil {
		return nil, err
	}
	return prs, nil
}

func (r *PullRequestRepo) GetByID(prID string) (*models.PullRequest, error) {
	var pr models.PullRequest
	if err := r.db.Where("id = ?", prID).First(&pr).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrPRNotFound
		}
		return nil, err
	}
	return &pr, nil
}

func (r *PullRequestRepo) Merge(prID string, merged_at time.Time) error {
	res := r.db.Model(&models.PullRequest{}).Where("id = ?", prID).Updates(map[string]interface{}{
		"status":    models.PRStatusMerged,
		"merged_at": &merged_at,
	})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return apperrors.ErrPRNotFound
	}
	return nil
}

func (r *PullRequestRepo) Reassign(prID, oldReviewerID, newReviewerID string) (*models.PullRequest, error) {
	var pr models.PullRequest
	if err := r.db.First(&pr, "id = ?", prID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrPRNotFound
		}
		return nil, err
	}
	for i, reviewer := range pr.Reviewers {
		if reviewer == oldReviewerID {
			pr.Reviewers[i] = newReviewerID
			break
		}
	}
	if err := r.db.Save(&pr).Error; err != nil {
		return nil, err
	}
	return &pr, nil
}
