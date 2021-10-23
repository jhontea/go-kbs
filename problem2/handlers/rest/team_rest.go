package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"go-kbs-soccer/entities/models"
	"go-kbs-soccer/entities/responses"
	"go-kbs-soccer/services"
)

type TeamRestHandler struct {
	service services.TeamService
}

func NewTeamRestHandler(service services.TeamService) *TeamRestHandler {
	return &TeamRestHandler{
		service: service,
	}
}

func (h *TeamRestHandler) GetTeam() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimSpace(r.URL.Query().Get("name"))
		if name == "" {
			responses.FailedResponse(w, nil, "team name must not empty")
			return
		}

		result, err := h.service.GetTeam(context.TODO(), name)
		if err != nil {
			responses.FailedResponse(w, nil, err.Error())
			return
		}

		responses.SuccessResponse(w, result, "success get team")
	}

}

func (h *TeamRestHandler) StoreTeam() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var team models.Team
		err := decoder.Decode(&team)
		if err != nil {
			responses.ErrorResponse(w, nil, "error payload", err.Error(), http.StatusBadRequest)
			return
		}

		if err = team.Validate(); err != nil {
			responses.FailedResponse(w, nil, err.Error())
			return
		}

		result, err := h.service.StoreTeam(context.TODO(), team)
		if err != nil {
			responses.FailedResponse(w, nil, err.Error())
			return
		}

		responses.SuccessResponse(w, result, "success store team")
	}

}
