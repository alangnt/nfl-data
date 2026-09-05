package utils

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func GetTeamColors(teamColors *[]TeamColor) (string, string) {
	var primaryColor string
	var secondaryColor string

	for _, color := range *teamColors {
		if color.Type == "primary" {
			primaryColor = color.HexColor
		}
		if color.Type == "secondary" {
			secondaryColor = color.HexColor
		}
	}

	return primaryColor, secondaryColor
}

func DisplayCard(primaryColor *string, secondaryColor *string, player *Player) string {
	jersey := player.Jersey
	if player.Jersey == "" {
		jersey = "No jersey number"
	}

	cardStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(*primaryColor)).
		BorderBackground(lipgloss.Color(*primaryColor)).
		Padding(1, 2).
		Margin(1, 0).
		Width(32)

	nameStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(*secondaryColor))

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#9CA3AF"))

	valStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(*secondaryColor)).
		Bold(true)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		nameStyle.Render(player.Name),
		"",
		lipgloss.JoinHorizontal(lipgloss.Left, labelStyle.Render("Jersey:   "), valStyle.Render(jersey)),
		lipgloss.JoinHorizontal(lipgloss.Left, labelStyle.Render("Position: "), valStyle.Render(player.Position)),
	)

	return cardStyle.Render(content)
}

func DisplayCards(primaryColor *string, secondaryColor *string, players *[]Player) {
	for _, player := range *players {
		fmt.Println(DisplayCard(primaryColor, secondaryColor, &player))
	}
}
