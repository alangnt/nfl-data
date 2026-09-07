package cards

import (
	"fmt"
	"nfl-data/types"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func DisplayGameCard(week *string, game *types.Game, isHomeTeam *bool) string {
	parsedTime, err := time.Parse(time.RFC3339, game.Scheduled)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	timeFormat := "Monday, Jan 2, 2006 at 3:04 PM"
	date := parsedTime.Format(timeFormat)

	home_points := game.Scoring.HomePoints
	away_points := game.Scoring.AwayPoints

	score := "Not played yet"
	if game.Status == "closed" {
		home_points_str := strconv.Itoa(home_points)
		away_points_str := strconv.Itoa(away_points)

		score = home_points_str + "-" + away_points_str
	}

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

	winStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#008000")).
		Bold(true)

	lossStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF0000")).
		Bold(true)

	var gameResult string
	if game.Status == "closed" {
		if (*isHomeTeam && home_points > away_points) || (!*isHomeTeam && away_points > home_points) {
			gameResult = winStyle.Render("W")
		} else {
			gameResult = lossStyle.Render("L")
		}
	} else {
		gameResult = ""
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		nameStyle.Render("Game", *week, gameResult),
		"",
		lipgloss.JoinHorizontal(lipgloss.Left, labelStyle.Render("Home: "), valStyle.Render(game.Home.Name)),
		lipgloss.JoinHorizontal(lipgloss.Left, labelStyle.Render("Away: "), valStyle.Render(game.Away.Name)),
		lipgloss.JoinHorizontal(lipgloss.Left, labelStyle.Render("Date: "), valStyle.Render(date)),
		"",
		lipgloss.JoinHorizontal(lipgloss.Left, labelStyle.Render("Score: "), valStyle.Render(score)),
	)

	return cardStyle.Render(content)
}

func DisplayGameCards(teamId *string, games *map[string]types.Game) {
	for week, game := range *games {
		isHomeTeam := false
		if game.Home.ID == *teamId {
			isHomeTeam = true
		}
		fmt.Println(DisplayGameCard(&week, &game, &isHomeTeam))
	}
}
