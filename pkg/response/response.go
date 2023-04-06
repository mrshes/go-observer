package response

import (
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func MakeResponse(w http.ResponseWriter, r *http.Request, status int, msg string, data interface{}) {
	response := Response{
		Status:  status,
		Message: msg,
		Data:    data,
	}
	render.Status(r, status)
	render.JSON(w, r, response)
}

func Success(w http.ResponseWriter, r *http.Request, msg string) {
	MakeResponse(w, r, http.StatusOK, msg, nil)
}

func JSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	MakeResponse(w, r, http.StatusOK, "SUCCESS", data)
}

func Error(w http.ResponseWriter, r *http.Request, error error) {
	MakeResponse(w, r, http.StatusBadRequest, error.Error(), nil)
}
