package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/MohdMusaiyab/cardyfy/models"
)

var githubAPI = "https://api.github.com"

// getGithubClient creates an HTTP client with Authorization if token is set
func getGithubClient() *http.Client {
	return &http.Client{}
}

// getAuthHeader returns Authorization header if GITHUB_TOKEN is set
func getAuthHeader() string {
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		return "token " + token
	}
	return ""
}

// FetchUser retrieves GitHub user profile
func FetchUser(username string) (*models.User, error) {
	url := fmt.Sprintf("%s/users/%s", githubAPI, username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if auth := getAuthHeader(); auth != "" {
		req.Header.Set("Authorization", auth)
	}

	client := getGithubClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user: %s", resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// FetchRepos retrieves public repositories of a user
func FetchRepos(username string) ([]models.Repo, error) {
	url := fmt.Sprintf("%s/users/%s/repos?per_page=100", githubAPI, username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if auth := getAuthHeader(); auth != "" {
		req.Header.Set("Authorization", auth)
	}

	client := getGithubClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch repos: %s", resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)

	var repos []models.Repo
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, err
	}

	return repos, nil
}

// FetchLanguages retrieves language breakdown for a repo
func FetchLanguages(username, repo string) (map[string]int, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/languages", githubAPI, username, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if auth := getAuthHeader(); auth != "" {
		req.Header.Set("Authorization", auth)
	}

	client := getGithubClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch languages: %s", resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)

	langs := make(map[string]int)
	if err := json.Unmarshal(body, &langs); err != nil {
		return nil, err
	}

	return langs, nil
}
