package server

import (
	"avitoTestTask/internal/infrastructure/transport/http/handlers"

	"github.com/go-chi/chi/v5"
)

func initRoutes(h *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/users/getReview", h.GetUserReviewPR)
	r.Post("/users/setIsActive", h.SetUserActive)

	r.Get("/team/get", h.GetTeamByName)
	r.Post("/team/add", h.CreateTeam)

	r.Post("/pullRequest/create", h.CreatePullRequest)
	r.Post("/pullRequest/merge", h.MergePullRequest)
	r.Post("/pullRequest/reassign", h.ReassignPullRequest)

	return r
}
