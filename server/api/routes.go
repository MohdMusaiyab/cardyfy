package api

import (
	"net/http"

	"github.com/MohdMusaiyab/cardyfy/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	// User profile
	mux.HandleFunc("/api/user/profile", handlers.GetUserProfile)

	// User repos
	mux.HandleFunc("/api/user/repos", handlers.GetUserRepos)

	// Repo languages
	mux.HandleFunc("/api/repo/languages", handlers.GetRepoLanguages)
}
