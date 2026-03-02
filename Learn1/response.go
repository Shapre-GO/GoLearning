package main

import (
	"errors"
	"net/http"
)

var errorStatusMap = map[error]int{
	ErrNameTooShort: http.StatusBadRequest,
	ErrInvalidEmail: http.StatusBadRequest,
	ErrEmailTaken:   http.StatusConflict,
}

func writeError(w http.ResponseWriter, err error) {
	for cause, status := range errorStatusMap {
		if errors.Is(err, cause) {
			http.Error(w, err.Error(), status)
			return
		}
	}

	http.Error(w, "SERVER ERROR", http.StatusInternalServerError)
}
