package models

// ThemeVariant represents the available card themes
type ThemeVariant string

// Define available theme constants
const (
    ThemeGithubDark  ThemeVariant = "github-dark"
    ThemeGithubLight ThemeVariant = "github-light"
    ThemeDracula     ThemeVariant = "dracula"
    ThemeNord        ThemeVariant = "nord"
    ThemeMonokai     ThemeVariant = "monokai"
    ThemeSolarized   ThemeVariant = "solarized"
    // Add more themes as needed
)

// GetThemeVariant converts a string to a ThemeVariant
func GetThemeVariant(theme string) ThemeVariant {
    switch theme {
    case string(ThemeGithubDark):
        return ThemeGithubDark
    case string(ThemeGithubLight):
        return ThemeGithubLight
    case string(ThemeDracula):
        return ThemeDracula
    case string(ThemeNord):
        return ThemeNord
    case string(ThemeMonokai):
        return ThemeMonokai
    case string(ThemeSolarized):
        return ThemeSolarized
    default:
        return ThemeGithubDark // Default theme
    }
}