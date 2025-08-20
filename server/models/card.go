package models

import "time"

// LangStat represents usage of a single language in repos.
type LangStat struct {
	Name     string  `json:"name"`
	Bytes    int     `json:"bytes"`
	Percent  float64 `json:"percent"`
}

// RepoBrief represents a minimal repo summary for cards.
type RepoBrief struct {
	Name      string   `json:"name"`
	Languages []string `json:"languages,omitempty"`
}

// CardData represents the extracted + aggregated profile data for rendering cards.
type CardData struct {
	// Identity
	Username    string  `json:"username"`
	DisplayName *string `json:"display_name,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
	Bio         *string `json:"bio,omitempty"`
	Location    *string `json:"location,omitempty"`
	Website     *string `json:"website,omitempty"`

	// Stats
	Followers   int `json:"followers"`
	Following   int `json:"following"`
	PublicRepos int `json:"public_repos"`

	TotalStars int `json:"total_stars"`
	TotalForks int `json:"total_forks"`

	// Languages
	PrimaryLanguage *string    `json:"primary_language,omitempty"`
	TopLanguages    []LangStat `json:"top_languages,omitempty"`

	// Repos
	TopReposByStars []RepoBrief `json:"top_repos_by_stars,omitempty"`
	TopRecentRepos  []RepoBrief `json:"top_recent_repos,omitempty"`

	// Timeline
	JoinedAt     *time.Time `json:"joined_at,omitempty"`
	LastActiveAt *time.Time `json:"last_active_at,omitempty"`

	// Extras
	Badges []string `json:"badges,omitempty"`
}
