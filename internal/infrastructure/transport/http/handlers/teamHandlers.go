package handlers

import (
	"avitoTestTask/internal/infrastructure/transport/http/dto"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
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

}

func decodeCreateTeamRequest(r *http.Request) (*dto.CreateTeamRequest, error) {
	var req dto.CreateTeamRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		return nil, err
	}
	return &req, nil
}
