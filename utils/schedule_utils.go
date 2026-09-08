package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"nfl-data/cards"
	"nfl-data/types"
)

func GetSchedule(teamID string, year string, seasonType string) {
	sportradarKey := GetSportradarAPIKey()

	url := "https://api.sportradar.com/nfl/official/trial/v7/en/games/" + year + "/" + seasonType + "/schedule.json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("x-api-key", sportradarKey)

	resp, err := sportradarClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Failed to close response body: %v", err)
		}
	}()

	var result types.Schedule
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	for _, week := range result.Weeks {
		for _, game := range week.Games {
			if game.Home.ID == teamID || game.Away.ID == teamID {
				isHomeTeam := false
				if game.Home.ID == teamID {
					isHomeTeam = true
				}

				fmt.Println(cards.DisplayGameCard(week.Sequence, game, isHomeTeam, seasonType))
			}
		}
	}
}
