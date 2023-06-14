package types

import (
	"github.com/go-chi/render"
	"net/http"
)

type ApiResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

func FailureResponse[T any](data T, writer http.ResponseWriter, request *http.Request) {
	anyResponse(false, data, writer, request)
}

func SuccessResponse[T any](data T, writer http.ResponseWriter, request *http.Request) {
	anyResponse(true, data, writer, request)
}

func anyResponse[T any](success bool, data T, writer http.ResponseWriter, request *http.Request) {
	response := ApiResponse[T]{
		Success: success,
		Data:    data,
	}

	render.JSON(writer, request, response)
}
