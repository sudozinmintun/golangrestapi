package routes

import (
	"net/http"

	"github.com/sudozinmintun/golangrestapi/internal/handlers"
)

func SetupHealthRoutes(mux http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("/health", handler.HealthHandler())
}
