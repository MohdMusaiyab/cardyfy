package models

// User maps to GitHub's /users/{username} response
type User struct {
	Login     string `json:"login"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Bio       string `json:"bio"`
	Company   string `json:"company"`
	Blog      string `json:"blog"`
	Location  string `json:"location"`
	Email     string `json:"email"`
	Twitter   string `json:"twitter_username"`

	PublicRepos int `json:"public_repos"`
	Followers   int `json:"followers"`
	Following   int `json:"following"`
}
