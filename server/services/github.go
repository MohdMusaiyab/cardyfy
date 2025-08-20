package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MohdMusaiyab/cardyfy/models"
)

func FetchUser(username string) (*models.GitHubUser, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s", username))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user models.GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func FetchRepos(username string) ([]models.GitHubRepo, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repos []models.GitHubRepo
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}
	return repos, nil
}

func FetchLanguages(username, repo string) (map[string]int, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/languages", username, repo))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	langs := make(map[string]int)
	if err := json.NewDecoder(resp.Body).Decode(&langs); err != nil {
		return nil, err
	}
	return langs, nil
}
