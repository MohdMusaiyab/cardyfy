package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MohdMusaiyab/cardyfy/models"
)

// GetUserRepos fetches all public repos of a GitHub user
func GetUserRepos(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username query parameter is required", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100", username)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch repositories from GitHub", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "GitHub API returned an error", resp.StatusCode)
		return
	}

	var repos []models.GitHubRepo
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		http.Error(w, "Failed to parse GitHub repos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repos)
}
