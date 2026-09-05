package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/charmbracelet/lipgloss"
)

type Player struct {
	Name     string `json:"name"`
	Jersey   string `json:"jersey"`
	Position string `json:"position"`
}

type TeamColor struct {
	Type     string `json:"type"`
	HexColor string `json:"hex_color"`
}

type Team struct {
	Name       string      `json:"name"`
	Market     string      `json:"market"`
	TeamColors []TeamColor `json:"team_colors"`
	Players    []Player    `json:"players"`
}

func GetTeam(teamId *string) {
	sportradarKey := GetSportraderAPIKey()

	url := "https://api.sportradar.com/nfl/official/trial/v7/en/teams/" + *teamId + "/profile.json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("x-api-key", sportradarKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Failed to close response body: %v", err)
		}
	}()

	var result Team
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	var primaryColor string
	var secondaryColor string

	for _, color := range result.TeamColors {
		if color.Type == "primary" {
			primaryColor = color.HexColor
		}
		if color.Type == "secondary" {
			secondaryColor = color.HexColor
		}
	}

	var positions []string

	fmt.Printf("Team: %s %s\n", result.Name, result.Market)
	for _, player := range result.Players {
		jersey := player.Jersey
		if player.Jersey == "" {
			jersey = "No jersey number"
		}

		positions = append(positions, player.Position)

		cardStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(primaryColor)).
			BorderBackground(lipgloss.Color(primaryColor)).
			Padding(1, 2).
			Margin(1, 0).
			Width(32)

		nameStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(secondaryColor))

		labelStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9CA3AF"))

		valStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color(secondaryColor)).
			Bold(true)

		content := lipgloss.JoinVertical(
			lipgloss.Left,
			nameStyle.Render(player.Name),
			"",
			lipgloss.JoinHorizontal(lipgloss.Left, labelStyle.Render("Jersey:   "), valStyle.Render(jersey)),
			lipgloss.JoinHorizontal(lipgloss.Left, labelStyle.Render("Position: "), valStyle.Render(player.Position)),
		)

		fmt.Println(cardStyle.Render(content))
	}

	slices.Sort(positions)
	for _, position := range slices.Compact(positions) {
		fmt.Println(position)
	}
}
