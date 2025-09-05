package handlers

import (
    "encoding/json"
    "net/http"
    "sync"

    "github.com/MohdMusaiyab/cardyfy/models"
    "github.com/MohdMusaiyab/cardyfy/services"
    "github.com/MohdMusaiyab/cardyfy/utils"
)

func GenerateCard(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")
    theme := r.URL.Query().Get("theme") // optional

    if username == "" {
        http.Error(w, "username query parameter is required", http.StatusBadRequest)
        return
    }

    // 1. Fetch user profile
    user, err := services.FetchUser(username)
    if err != nil {
        http.Error(w, "Failed to fetch GitHub user profile", http.StatusInternalServerError)
        return
    }

    // 2. Fetch repos
    repos, err := services.FetchRepos(username)
    if err != nil {
        http.Error(w, "Failed to fetch GitHub repos", http.StatusInternalServerError)
        return
    }

    // 3. Fetch languages concurrently
    languages := make(map[string]int)
    var mu sync.Mutex
    var wg sync.WaitGroup

    for _, repo := range repos {
        wg.Add(1)
        go func(repoName string) {
            defer wg.Done()
            repoLangs, err := services.FetchLanguages(username, repoName)
            if err != nil {
                return
            }
            mu.Lock()
            for lang, count := range repoLangs {
                languages[lang] += count
            }
            mu.Unlock()
        }(repo.Name)
    }
    wg.Wait()

    // 4. Build tech stack
    techStack := make([]string, 0, len(languages))
    for lang := range languages {
        techStack = append(techStack, lang)
    }

    // 5. Pick theme + variant (randomness handled inside utils)
    // Fix: Pass theme directly instead of in a slice
    chosenTheme := utils.PickThemeVariant(theme)

    // 6. Generate HTML card using service
    cardHTML, err := services.GenerateCardHTML(user, repos, techStack, models.ThemeVariant(chosenTheme))
    if err != nil {
        http.Error(w, "Failed to generate card", http.StatusInternalServerError)
        return
    }

    // 👉 For now: return JSON with HTML as preview (later: render to image)
    response := map[string]interface{}{
        "user":   user,
        "theme":  chosenTheme,
        "card":   cardHTML,
        "status": "success",
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}