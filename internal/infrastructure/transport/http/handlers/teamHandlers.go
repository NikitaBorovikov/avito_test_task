package handlers

import (
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
		sendErrorResponse(w, r, http.StatusBadRequest, "NOT_FOUND", "Invalid request format")
		return
	}

	team := req.ToDomainTeam()

	res, err := h.TeamUC.Create(&team)
	if err != nil {
		// TODO: добавить обработку ошибок
		logrus.Errorf("failed to create team: %v", err)
		return
	}

	resp := dto.NewCreateTeamOKResponse(res)
	sendOkResponse(w, r, http.StatusCreated, resp)
}

func (h *Handlers) GetTeamByName(w http.ResponseWriter, r *http.Request) {
	teamName := r.URL.Query().Get(getTeamQueryParam)

	team, err := h.TeamUC.GetByName(teamName)
	if err != nil {
		// TODO: handle err
		logrus.Errorf("failed to get team: %v", err)
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
