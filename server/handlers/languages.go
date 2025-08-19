package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MohdMusaiyab/cardyfy/models"
)

// GetRepoLanguages fetches languages used in a given repo
func GetRepoLanguages(w http.ResponseWriter, r *http.Request) {
	owner := r.URL.Query().Get("owner")
	repo := r.URL.Query().Get("repo")

	if owner == "" || repo == "" {
		http.Error(w, "owner and repo query parameters are required", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/languages", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch languages from GitHub", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "GitHub API returned an error", resp.StatusCode)
		return
	}

	var languages models.GitHubLanguages
	if err := json.NewDecoder(resp.Body).Decode(&languages); err != nil {
		http.Error(w, "Failed to parse GitHub languages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(languages)
}
