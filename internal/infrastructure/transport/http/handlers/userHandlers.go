package handlers

import (
	"avitoTestTask/internal/infrastructure/transport/http/dto"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

const (
	userIDQueryParam = "user_id"
)

func (h *Handlers) GetUserReviewPR(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get(userIDQueryParam)

	if err := dto.ValidateUserID(userID); err != nil {
		logrus.Errorf("validate error: %v", err)
		sendErrorResponse(w, r, http.StatusBadRequest, "NOT_FOUND", err.Error())
		return
	}

	prs, err := h.PullRequestUC.GetByReviewer(userID)
	if err != nil {
		// TODO: handle error
		return
	}

	resp := dto.NewGetUserReviewPRResponse(userID, prs)
	sendOkResponse(w, r, http.StatusOK, resp)
}

func (h *Handlers) SetUserActive(w http.ResponseWriter, r *http.Request) {
	req, err := decodeSetUserActiveReq(r)
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

	user, err := h.UserUC.SetUserActive(req.UserID, req.IsActive)
	if err != nil {
		// TODO: handle errors
		return
	}

	// Get team name from DB
	teamName := "backend"
	resp := dto.NewSetUserActiveResponse(teamName, user)
	sendOkResponse(w, r, http.StatusOK, resp)
}

func decodeSetUserActiveReq(r *http.Request) (*dto.SetUserActiveRequest, error) {
	var req dto.SetUserActiveRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		return nil, err
	}
	return &req, nil
}
