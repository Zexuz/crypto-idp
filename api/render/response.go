package render

import (
	"github.com/go-chi/render"
	"net/http"
)

func GenericError(writer http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusInternalServerError)
	render.JSON(writer, request, map[string]string{
		"error": "An error occurred",
	})
}

func Error(writer http.ResponseWriter, request *http.Request, err error, status int) {
	render.Status(request, status)
	render.JSON(writer, request, map[string]string{
		"error": err.Error(),
	})
}
func OK[T any](writer http.ResponseWriter, request *http.Request, data T) {
	render.Status(request, http.StatusOK)
	render.JSON(writer, request, data)
}
