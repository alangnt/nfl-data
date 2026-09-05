package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"

	"nfl-data/cards"
	"nfl-data/types"
)

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

	var result types.Team
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	primaryColor, secondaryColor := cards.GetTeamColors(&result.TeamColors)

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

		var players []types.Player
		for _, player := range result.Players {
			if player.Position == position {
				players = append(players, player)
			}
		}

		cards.DisplayPlayerCards(&primaryColor, &secondaryColor, &players)
	}
}
