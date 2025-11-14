package handlers

import (
	"avitoTestTask/internal/infrastructure/transport/http/dto"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

func (h *Handlers) CreatePullRequest(w http.ResponseWriter, r *http.Request) {
	req, err := decodeCreatePullRequest(r)
	if err != nil {
		logrus.Errorf("failed to decode JSON: %v", err)
		sendErrorResponse(w, r, http.StatusBadRequest, "NOT_FOUND", "Invalid request format")
		return
	}

	pr := req.ToDomainPR()

	pullRequest, err := h.PullRequestUC.Create(&pr)
	if err != nil {
		// TODO: handle errors
		return
	}

	resp := dto.NewCreatePRResponse(pullRequest)
	sendOkResponse(w, r, http.StatusCreated, resp)
}

func (h *Handlers) MergePullRequest(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) ReassignPullRequest(w http.ResponseWriter, r *http.Request) {

}

func decodeCreatePullRequest(r *http.Request) (*dto.CreatePRRequest, error) {
	var req dto.CreatePRRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		return nil, err
	}
	return &req, nil
}
