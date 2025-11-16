package repository

import "avitoTestTask/internal/core/models"

type StatsRepo interface {
	GetReviewerStats() ([]models.ReviewerStats, error)
}
