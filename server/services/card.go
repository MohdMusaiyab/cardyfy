package services

import (
	"sort"
	"time"

	"github.com/MohdMusaiyab/cardyfy/models"
)

// BuildCardData transforms UserProfileResponse into CardData for rendering.
func BuildCardData(resp models.UserProfileResponse) models.CardData {
	// -------------------------------
	// Identity
	// -------------------------------
	card := models.CardData{
		Username:    resp.User.Login,
		DisplayName: &resp.User.Name,
		AvatarURL:   &resp.User.AvatarURL,
		Bio:         &resp.User.Bio,
		Location:    &resp.User.Location,
		Website:     &resp.User.Blog,
		Followers:   resp.User.Followers,
		Following:   resp.User.Following,
		PublicRepos: resp.User.PublicRepos,
		JoinedAt:    parseTimePtr(resp.User.CreatedAt),
	}

	// -------------------------------
	// Languages (aggregate)
	// -------------------------------
	var totalBytes int
	for _, count := range resp.Languages {
		totalBytes += count
	}
	langStats := make([]models.LangStat, 0, len(resp.Languages))
	for lang, count := range resp.Languages {
		percent := 0.0
		if totalBytes > 0 {
			percent = (float64(count) / float64(totalBytes)) * 100
		}
		langStats = append(langStats, models.LangStat{
			Name:    lang,
			Bytes:   count,
			Percent: percent,
		})
	}
	// Sort languages by bytes desc
	sort.Slice(langStats, func(i, j int) bool {
		return langStats[i].Bytes > langStats[j].Bytes
	})
	card.TopLanguages = langStats
	if len(langStats) > 0 {
		card.PrimaryLanguage = &langStats[0].Name
	}

	// -------------------------------
	// Repos: stars + recent
	// -------------------------------
	// sort repos by stars
	repos := resp.Repos
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].Stargazers > repos[j].Stargazers
	})
	topStarRepos := []models.RepoBrief{}
	for i, repo := range repos {
		if i >= 5 {
			break
		}
		topStarRepos = append(topStarRepos, models.RepoBrief{
			Name:      repo.Name,
			Languages: []string{repo.Language}, // simplified
		})
	}
	card.TopReposByStars = topStarRepos

	// sort repos by updated_at
	sort.Slice(repos, func(i, j int) bool {
		ti, erri := time.Parse(time.RFC3339, repos[i].UpdatedAt)
		tj, errj := time.Parse(time.RFC3339, repos[j].UpdatedAt)
		if erri != nil && errj != nil {
			return false
		}
		if erri != nil {
			return false
		}
		if errj != nil {
			return true
		}
		return ti.After(tj)
	})
	topRecentRepos := []models.RepoBrief{}
	for i, repo := range repos {
		if i >= 5 {
			break
		}
		topRecentRepos = append(topRecentRepos, models.RepoBrief{
			Name:      repo.Name,
			Languages: []string{repo.Language},
		})
	}
	card.TopRecentRepos = topRecentRepos

	// -------------------------------
	// Total stars/forks
	// -------------------------------
	totalStars := 0
	totalForks := 0
	for _, repo := range repos {
		totalStars += repo.Stargazers
		totalForks += repo.Forks
	}
	card.TotalStars = totalStars
	card.TotalForks = totalForks

	// -------------------------------
	// Last active (latest repo update)
	// -------------------------------
	if len(repos) > 0 {
		last := repos[0].UpdatedAt
		lastTime, err := time.Parse(time.RFC3339, last)
		if err == nil {
			for _, repo := range repos {
				repoTime, err := time.Parse(time.RFC3339, repo.UpdatedAt)
				if err == nil && repoTime.After(lastTime) {
					lastTime = repoTime
				}
			}
		}
		card.LastActiveAt = &lastTime
	}

	return card
}

func parseTimePtr(s string) *time.Time {
	if s == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil
	}
	return &t
}
