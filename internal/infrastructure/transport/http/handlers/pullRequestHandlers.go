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

	if err := req.Validate(); err != nil {
		logrus.Errorf("validate error: %v", err)
		sendErrorResponse(w, r, http.StatusBadRequest, "NOT_FOUND", err.Error())
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
	req, err := decodeMergePRRequest(r)
	if err != nil {
		logrus.Errorf("failed to decode JSON: %v", err)
		sendErrorResponse(w, r, http.StatusBadRequest, "NOT_FOUND", "Invalid request format")
		return
	}

	if err := req.Validate(); err != nil {
		logrus.Errorf("validate error: %v", err)
		sendErrorResponse(w, r, http.StatusBadRequest, "NOT_FOUND", err.Error())
		return
	}

	pullRequest, err := h.PullRequestUC.Merge(req.PullRequestID)
	if err != nil {
		// TODO: handle errors
		return
	}

	resp := dto.NewMergePRResponse(pullRequest)
	sendOkResponse(w, r, http.StatusOK, resp)
}

func (h *Handlers) ReassignPullRequest(w http.ResponseWriter, r *http.Request) {
	req, err := decodeReassignPRRequest(r)
	if err != nil {
		logrus.Errorf("failed to decode JSON: %v", err)
		sendErrorResponse(w, r, http.StatusBadRequest, "NOT_FOUND", "Invalid request format")
		return
	}

	if err := req.Validate(); err != nil {
		logrus.Errorf("validate error: %v", err)
		sendErrorResponse(w, r, http.StatusBadRequest, "NOT_FOUND", err.Error())
		return
	}

	pullRequest, err := h.PullRequestUC.Reassign(req.PullRequestID, req.OldUserID)
	if err != nil {
		// TODO: handle errors
		return
	}

	resp := dto.NewReassignPRResponse(pullRequest)
	sendOkResponse(w, r, http.StatusOK, resp)
}

func decodeCreatePullRequest(r *http.Request) (*dto.CreatePRRequest, error) {
	var req dto.CreatePRRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func decodeMergePRRequest(r *http.Request) (*dto.MergePRRequest, error) {
	var req dto.MergePRRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func decodeReassignPRRequest(r *http.Request) (*dto.ReassignPRRequest, error) {
	var req dto.ReassignPRRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		return nil, err
	}
	return &req, nil
}
