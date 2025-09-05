package models

// UserProfileResponse is our custom aggregated response
type UserProfileResponse struct {
	User      User            `json:"user"`
	Repos     []Repo          `json:"repos"`
	Languages map[string]int  `json:"languages"`
	TechStack []string        `json:"tech_stack"`
}
