package handlers

import (
	apperrors "avitoTestTask/internal/appErrors"
	"avitoTestTask/internal/infrastructure/transport/http/dto"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

const (
	getTeamQueryParam = "team_name"
)

func (h *Handlers) CreateTeam(w http.ResponseWriter, r *http.Request) {
	req, err := decodeCreateTeamRequest(r)
	if err != nil {
		logrus.Errorf("failed to decode JSON: %v", err)
		sendErrorResponse(w, r, http.StatusBadRequest, apperrors.ErrorCodeNotFound, "Invalid request format")
		return
	}

	if err := req.Validate(); err != nil {
		logrus.Errorf("validate error: %v", err)
		sendErrorResponse(w, r, http.StatusBadRequest, apperrors.ErrorCodeNotFound, err.Error())
		return
	}

	team := req.ToDomainTeam()
	res, err := h.TeamUC.Create(team)
	if err != nil {
		logrus.Errorf("failed to create team: %v", err)
		errInfo := apperrors.HandleError(err)
		sendErrorResponse(w, r, errInfo.HttdCode, errInfo.Code, errInfo.Msg)
		return
	}

	resp := dto.NewCreateTeamOKResponse(res)
	sendOkResponse(w, r, http.StatusCreated, resp)
}

func (h *Handlers) GetTeamByName(w http.ResponseWriter, r *http.Request) {
	teamName := r.URL.Query().Get(getTeamQueryParam)

	if err := dto.ValidateTeamName(teamName); err != nil {
		logrus.Errorf("validate error: %v", err)
		sendErrorResponse(w, r, http.StatusBadRequest, apperrors.ErrorCodeNotFound, err.Error())
		return
	}

	team, err := h.TeamUC.GetByName(teamName)
	if err != nil {
		logrus.Errorf("failed to get team: %v", err)
		errInfo := apperrors.HandleError(err)
		sendErrorResponse(w, r, errInfo.HttdCode, errInfo.Code, errInfo.Msg)
		return
	}

	resp := dto.NewGetTeamByNameResponse(team)
	sendOkResponse(w, r, http.StatusOK, resp)
}

func decodeCreateTeamRequest(r *http.Request) (*dto.CreateTeamRequest, error) {
	var req dto.CreateTeamRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		return nil, err
	}
	return &req, nil
}
