package responses

import (
	"encoding/json"
	"net/http"
)

type RestResponse struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

type ResponseBuilder struct {
	Writer     http.ResponseWriter
	StatusCode int
	Message    string      `json:"message"`
	Error      string      `json:"error"`
	Data       interface{} `json:"data"`
}

func Response(w http.ResponseWriter) *ResponseBuilder {
	return &ResponseBuilder{
		Writer: w,
	}
}

func (r *ResponseBuilder) SetStatus(statusCode int) *ResponseBuilder {
	r.StatusCode = statusCode
	return r
}

func (r *ResponseBuilder) SetMessage(message string) *ResponseBuilder {
	r.Message = message
	return r
}

func (r *ResponseBuilder) SetData(data interface{}) *ResponseBuilder {
	r.Data = data
	return r
}

func (r *ResponseBuilder) SetError(err string) *ResponseBuilder {
	r.Error = err
	return r
}

func (r *ResponseBuilder) Build() {
	r.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(RestResponse{
		Message: r.Message,
		Data:    r.Data,
		Error:   r.Error,
	})

	r.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.Writer.WriteHeader(r.StatusCode)
	r.Writer.Write(result)
}

func SuccessResponse(w http.ResponseWriter, data interface{}, message string) {
	Response(w).
		SetData(data).
		SetMessage(message).
		SetStatus(http.StatusOK).
		Build()
}

func FailedResponse(w http.ResponseWriter, data interface{}, message string) {
	Response(w).
		SetData(data).
		SetMessage(message).
		SetStatus(http.StatusBadRequest).
		Build()
}

func ErrorResponse(w http.ResponseWriter, data interface{}, message, err string, statusCode int) {
	Response(w).
		SetData(data).
		SetMessage(message).
		SetError(err).
		SetStatus(statusCode).
		Build()
}

func InternalErrorResponse(w http.ResponseWriter, message string) {
	Response(w).
		SetMessage(message).
		SetStatus(http.StatusInternalServerError).
		Build()
}
