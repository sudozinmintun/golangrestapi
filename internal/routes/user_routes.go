package routes

import (
	"net/http"

	"github.com/sudozinmintun/golangrestapi/internal/handlers"
)

func SetupUserRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("POST /user/register", handler.CreateUserHandler())
}
