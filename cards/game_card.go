package cards

import (
	"fmt"
	"nfl-data/types"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func DisplayGameCard(week *string, game *types.Game) string {
	parsedTime, err := time.Parse(time.RFC3339, game.Scheduled)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	timeFormat := "Monday, Jan 2, 2006 at 3:04 PM"
	date := parsedTime.Format(timeFormat)

	cardStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FFFFFF")).
		BorderBackground(lipgloss.Color("#FFFFFF")).
		Padding(1, 2).
		Margin(1, 0).
		Width(32)

	nameStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF"))

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#9CA3AF"))

	valStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		nameStyle.Render("Game", *week),
		"",
		lipgloss.JoinHorizontal(lipgloss.Left, labelStyle.Render("Home: "), valStyle.Render(game.Home.Name)),
		lipgloss.JoinHorizontal(lipgloss.Left, labelStyle.Render("Away: "), valStyle.Render(game.Away.Name)),
		lipgloss.JoinHorizontal(lipgloss.Left, labelStyle.Render("Date: "), valStyle.Render(date)),
	)

	return cardStyle.Render(content)
}

func DisplayGameCards(games *map[string]types.Game) {
	for week, game := range *games {
		fmt.Println(DisplayGameCard(&week, &game))
	}
}
