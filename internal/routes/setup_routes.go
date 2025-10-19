package routes

import (
	"net/http"

	"github.com/sudozinmintun/golangrestapi/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	SetupHealthRoutes(*mux, handler)
}
