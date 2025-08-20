package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MohdMusaiyab/cardyfy/services"
)

func GetUserLanguages(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	repo := r.URL.Query().Get("repo")

	if username == "" || repo == "" {
		http.Error(w, "username and repo query parameters are required", http.StatusBadRequest)
		return
	}

	langs, err := services.FetchLanguages(username, repo)
	if err != nil {
		http.Error(w, "Failed to fetch languages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(langs)
}
