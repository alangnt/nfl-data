package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"nfl-data/cards"
	"nfl-data/types"
)

func GetSchedule(teamId *string) {
	sportradarKey := GetSportraderAPIKey()

	url := "https://api.sportradar.com/nfl/official/trial/v7/en/games/2026/REG/schedule.json"

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

	var result types.Schedule
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	games := make(map[string]types.Game)

	for _, week := range result.Weeks {
		for _, game := range week.Games {
			if game.Home.ID == *teamId || game.Away.ID == *teamId {
				games[week.Title] = game
			}
		}
	}

	cards.DisplayGameCards(&games)
}
