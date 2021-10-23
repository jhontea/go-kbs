package rest

import (
	"encoding/json"
	"go-kbs-notification/entities/models"
	"go-kbs-notification/entities/responses"
	"go-kbs-notification/services"
	"net/http"
)

type FeatureFlagRestHandler struct {
	service services.FeatureFlagService
}

func NewFeatureFlagRestHandler(service services.FeatureFlagService) *FeatureFlagRestHandler {
	return &FeatureFlagRestHandler{
		service: service,
	}
}

func (h *FeatureFlagRestHandler) GetFeatureFlag() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		result, err := h.service.GetFeatureFlag()
		if err != nil {
			responses.FailedResponse(w, nil, err.Error())
			return
		}

		responses.SuccessResponse(w, result, "success get feature flag")
	}

}

func (h *FeatureFlagRestHandler) StoreFeatureFlag() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var featureFlag []models.FeatureFlag
		err := decoder.Decode(&featureFlag)
		if err != nil {
			responses.ErrorResponse(w, nil, "error payload", err.Error(), http.StatusBadRequest)
			return
		}

		err = h.service.StoreFeatureFlag(featureFlag)
		if err != nil {
			responses.FailedResponse(w, nil, err.Error())
			return
		}

		responses.SuccessResponse(w, nil, "success store feture flag")
	}

}
