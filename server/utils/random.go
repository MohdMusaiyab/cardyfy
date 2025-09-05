package utils

import (
    "math/rand"
    "strings"
    "time"
)

// PickThemeVariant selects a theme variant based on user's preference or randomly
func PickThemeVariant(themePreference string) string {
    // List of available themes
    themes := []string{
        "github-dark",
        "github-light",
        "dracula",
        "nord",
        "monokai",
        "solarized",
    }
    
    // If a specific theme is requested and it's valid
    if themePreference != "" {
        themePreference = strings.ToLower(themePreference)
        for _, t := range themes {
            if t == themePreference {
                return t
            }
        }
    }
    
    // If no theme specified or invalid theme, pick random
    rand.Seed(time.Now().UnixNano())
    return themes[rand.Intn(len(themes))]
}