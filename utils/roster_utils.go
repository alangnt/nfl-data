package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

type Player struct {
	Name     string `json:"name"`
	Jersey   string `json:"jersey"`
	Position string `json:"position"`
}

type Roster struct {
	Name    string   `json:"name"`
	Market  string   `json:"market"`
	Players []Player `json:"players"`
}

func GetTeam(teamId *string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sportradar := os.Getenv("SPORTRADAR_API_KEY")
	if sportradar == "" {
		log.Fatal("API key required")
	}

	url := "https://api.sportradar.com/nfl/official/trial/v7/en/teams/" + *teamId + "/full_roster.json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("x-api-key", sportradar)

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

	var result Roster
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	fmt.Printf("Team: %s %s\n", result.Name, result.Market)
	for _, player := range result.Players {
		jersey := player.Jersey
		if player.Jersey == "" {
			jersey = "No jersey number"
		}

		cardStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#E31837")). // Chiefs Red
			Padding(1, 2).
			Width(32)

		nameStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF"))

		labelStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9CA3AF"))

		valStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F3F4F6")).
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
}
