package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
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

	primaryColor, secondaryColor := GetTeamColors(&result.TeamColors)

	var uniquePositions []string

	fmt.Printf("Team: %s %s\n", result.Name, result.Market)
	for _, player := range result.Players {
		uniquePositions = append(uniquePositions, player.Position)
	}

	slices.Sort(uniquePositions)
	positions := slices.Compact(uniquePositions)

	for {
		position := GetPosition(&positions)

		if position == "exit" {
			break
		}

		var players []Player
		for _, player := range result.Players {
			if player.Position == position {
				players = append(players, player)
			}
		}

		DisplayCards(&primaryColor, &secondaryColor, &players)
	}
}
