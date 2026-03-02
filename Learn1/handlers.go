package main

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	repo UserRepository
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Failed to decode response", http.StatusBadRequest)
		return
	}

	if err := u.Validate(); err != nil {
		writeError(w, err)
		return
	}

	if err := h.repo.Create(&u); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
