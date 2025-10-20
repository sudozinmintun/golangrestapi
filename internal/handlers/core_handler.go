package handlers

import (
	"database/sql"

	"github.com/sudozinmintun/golangrestapi/internal/store"
)

type Handler struct {
	DB      *sql.DB
	Queries *store.Queries
}

func NewHandler(db *sql.DB, queries *store.Queries) *Handler {
	return &Handler{
		DB:      db,
		Queries: queries,
	}
}
