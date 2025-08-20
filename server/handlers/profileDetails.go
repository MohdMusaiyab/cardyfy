package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/MohdMusaiyab/cardyfy/models"
	"github.com/MohdMusaiyab/cardyfy/services"
)

func GetUserProfileSummary(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username query parameter is required", http.StatusBadRequest)
		return
	}

	// -----------------------------
	// 1. Fetch user profile
	// -----------------------------
	user, err := services.FetchUser(username)
	if err != nil {
		http.Error(w, "Failed to fetch GitHub user profile", http.StatusInternalServerError)
		return
	}

	// -----------------------------
	// 2. Fetch user repositories
	// -----------------------------
	repos, err := services.FetchRepos(username)
	if err != nil {
		http.Error(w, "Failed to fetch GitHub repos", http.StatusInternalServerError)
		return
	}

	// -----------------------------
	// 3. Fetch languages concurrently
	// -----------------------------
	languages := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, repo := range repos {
		wg.Add(1)
		go func(repoName string) {
			defer wg.Done()

			repoLangs, err := services.FetchLanguages(username, repoName)
			if err != nil {
				return // skip failed repo
			}

			mu.Lock()
			for lang, count := range repoLangs {
				languages[lang] += count
			}
			mu.Unlock()
		}(repo.Name)
	}

	wg.Wait()

	// -----------------------------
	// 4. Build TechStack (unique langs)
	// -----------------------------
	techStack := make([]string, 0, len(languages))
	for lang := range languages {
		techStack = append(techStack, lang)
	}

	// -----------------------------
	// 5. Final aggregated response
	// -----------------------------
	response := models.UserProfileResponse{
		User:      *user,
		Repos:     repos,
		Languages: languages,
		TechStack: techStack,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


