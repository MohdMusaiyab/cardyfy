package services

import (
    "bytes"
    "fmt"
    "html/template"
    "path/filepath"

    "github.com/MohdMusaiyab/cardyfy/models"
)

// GenerateCardHTML renders a user card using HTML templates + theme
func GenerateCardHTML(
    user *models.User,
    repos []models.Repo,
    techStack []string,
    theme models.ThemeVariant,
) (string, error) {

    // Build the path for the selected theme
    // Fix: Use string(theme) instead of theme.Type
    themeFile := filepath.Join("templates", "themes", fmt.Sprintf("%s.html", string(theme)))

    // Parse both base.html and the theme file
    tmpl, err := template.ParseFiles("templates/base.html", themeFile)
    if err != nil {
        return "", err
    }

    // Data passed into template
    data := struct {
        User      *models.User
        Repos     []models.Repo
        TechStack []string
        Theme     models.ThemeVariant
    }{
        User:      user,
        Repos:     repos,
        TechStack: techStack,
        Theme:     theme,
    }

    // Render template into buffer
    var buf bytes.Buffer
    err = tmpl.Execute(&buf, data)
    if err != nil {
        return "", err
    }

    return buf.String(), nil
}