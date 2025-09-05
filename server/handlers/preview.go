package handlers

import (
    "net/http"
    "sync"

    "github.com/MohdMusaiyab/cardyfy/models"
    "github.com/MohdMusaiyab/cardyfy/services"
    "github.com/MohdMusaiyab/cardyfy/utils"
)

func PreviewCard(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")
    theme := r.URL.Query().Get("theme")

    if username == "" {
        http.Error(w, "username query parameter is required", http.StatusBadRequest)
        return
    }

    // 1. Fetch GitHub profile + repos + langs
    user, _ := services.FetchUser(username)
    repos, _ := services.FetchRepos(username)

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

    techStack := make([]string, 0, len(languages))
    for lang := range languages {
        techStack = append(techStack, lang)
    }

    // 2. Pick theme
    chosenTheme := utils.PickThemeVariant(theme)

    // 3. Generate HTML card
    cardHTML, err := services.GenerateCardHTML(user, repos, techStack, models.ThemeVariant(chosenTheme))
    if err != nil {
        http.Error(w, "failed to render card", http.StatusInternalServerError)
        return
    }

    // 4. Serve as HTML
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(cardHTML))
}