package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MohdMusaiyab/cardyfy/services"
)

func GetUserRepos(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username query parameter is required", http.StatusBadRequest)
		return
	}

	repos, err := services.FetchRepos(username)
	if err != nil {
		http.Error(w, "Failed to fetch repos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repos)
}
