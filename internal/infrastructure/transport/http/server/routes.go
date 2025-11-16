package server

import (
	"avitoTestTask/internal/infrastructure/transport/http/handlers"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func initRoutes(h *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/users/getReview", h.GetUserReviewPR)
	r.Post("/users/setIsActive", h.SetUserActive)

	r.Get("/team/get", h.GetTeamByName)
	r.Post("/team/add", h.CreateTeam)

	r.Post("/pullRequest/create", h.CreatePullRequest)
	r.Post("/pullRequest/merge", h.MergePullRequest)
	r.Post("/pullRequest/reassign", h.ReassignPullRequest)

	r.Get("/stats/reviewers", h.GetReviewerStats)

	return r
}
