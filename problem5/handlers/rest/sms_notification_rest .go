package rest

import (
	"encoding/json"
	"go-kbs-notification/entities/models"
	"go-kbs-notification/entities/responses"
	"go-kbs-notification/services"
	"net/http"
)

type SMSNotificationRestHandler struct {
	service services.SMSNotificationService
}

func NewSMSNotificationRestHandler(service services.SMSNotificationService) *SMSNotificationRestHandler {
	return &SMSNotificationRestHandler{
		service: service,
	}
}

func (h *SMSNotificationRestHandler) SendSMS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var req models.SMS
		err := decoder.Decode(&req)
		if err != nil {
			responses.ErrorResponse(w, nil, "error payload", err.Error(), http.StatusBadRequest)
			return
		}

		err = h.service.SendSMS(req)
		if err != nil {
			responses.FailedResponse(w, nil, err.Error())
			return
		}

		responses.SuccessResponse(w, nil, "success send sms")
	}

}
