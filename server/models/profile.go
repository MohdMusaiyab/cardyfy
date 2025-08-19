package models

// UserProfileResponse is the final aggregated response for the frontend
type UserProfileResponse struct {
	User      GitHubUser              `json:"user"`
	Repos     []GitHubRepo            `json:"repos"`
	Languages map[string]int          `json:"languages"` // aggregated languages across repos
	TechStack []string                `json:"tech_stack"` // simplified list of top languages/tech
}
