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

type PlayerRestHandler struct {
	service services.PlayerService
}

func NewPlayerRestHandler(service services.PlayerService) *PlayerRestHandler {
	return &PlayerRestHandler{
		service: service,
	}
}

func (h *PlayerRestHandler) GetPlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimSpace(r.URL.Query().Get("name"))
		if name == "" {
			responses.FailedResponse(w, nil, "player name must not empty")
			return
		}

		result, err := h.service.GetPlayer(context.TODO(), name)
		if err != nil {
			responses.FailedResponse(w, nil, err.Error())
			return
		}

		responses.SuccessResponse(w, result, "success get player")
	}

}

func (h *PlayerRestHandler) StorePlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var player models.Player
		err := decoder.Decode(&player)
		if err != nil {
			responses.ErrorResponse(w, nil, "error payload", err.Error(), http.StatusBadRequest)
			return
		}

		if err = player.Validate(); err != nil {
			responses.FailedResponse(w, nil, err.Error())
			return
		}

		result, err := h.service.StorePlayer(context.TODO(), player)
		if err != nil {
			responses.FailedResponse(w, nil, err.Error())
			return
		}

		responses.SuccessResponse(w, result, "success store player")
	}

}
