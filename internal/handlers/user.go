package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sudozinmintun/golangrestapi/internal/dtos"
	"github.com/sudozinmintun/golangrestapi/internal/store"
	"github.com/sudozinmintun/golangrestapi/internal/utils"
)

func (h *Handler) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dtos.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.ResposnseWithError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			utils.ResposnseWithError(w, http.StatusInternalServerError, "failed to hash password")
			return
		}

		_, err = h.Queries.CreateUser(ctx, store.CreateUserParams{
			Username: req.Username,
			Email:    req.Email,
			Password: hashedPassword,
		})

		if err != nil {
			utils.ResposnseWithError(w, http.StatusInternalServerError, "failed to create user")
			return
		}

		utils.ResponseWithSuccess(w, http.StatusCreated, "user created", req.Username)
	}
}
