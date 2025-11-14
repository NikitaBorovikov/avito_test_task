package handlers

import (
	"avitoTestTask/internal/infrastructure/transport/http/dto"
	"avitoTestTask/internal/usecases"
	"net/http"

	"github.com/go-chi/render"
)

type Handlers struct {
	UserUC        *usecases.UserUC
	TeamUC        *usecases.TeamUC
	PullRequestUC *usecases.PullRequestUC
}

func NewHandlers(uc *usecases.UseCases) *Handlers {
	return &Handlers{
		UserUC:        uc.UserUC,
		TeamUC:        uc.TeamUC,
		PullRequestUC: uc.PullRequestUC,
	}
}

func sendOkResponse(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	render.JSON(w, r, data)
}

func sendErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, code, msg string) {
	w.WriteHeader(statusCode)
	render.JSON(w, r, dto.NewErrorResponse(code, msg))
}
