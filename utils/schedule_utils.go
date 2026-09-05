package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type GameTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Scoring struct {
	HomePoints int `json:"home_points"`
	AwayPoints int `json:"away_points"`
}

type Game struct {
	Home      GameTeam `json:"home"`
	Away      GameTeam `json:"away"`
	Scoring   Scoring  `json:"scoring"`
	Status    string   `json:"status"`
	Scheduled string   `json:"scheduled"`
}

type Week struct {
	Games []Game `json:"games"`
	Title string `json:"title"`
}

type Schedule struct {
	Weeks []Week `json:"weeks"`
}

func GetSchedule(teamId *string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sportradar := os.Getenv("SPORTRADAR_API_KEY")
	if sportradar == "" {
		log.Fatal("API key required")
	}

	url := "https://api.sportradar.com/nfl/official/trial/v7/en/games/2026/REG/schedule.json"

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

	var result Schedule
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	for _, week := range result.Weeks {
		for _, game := range week.Games {
			if game.Home.ID == *teamId || game.Away.ID == *teamId {
				parsedTime, err := time.Parse(time.RFC3339, game.Scheduled)
				if err != nil {
					fmt.Println("Error parsing date:", err)
					return
				}

				timeFormat := "Monday, Jan 2, 2006 at 3:04 PM"
				date := parsedTime.Format(timeFormat)

				fmt.Printf("Week %s\n", week.Title)
				fmt.Println(game.Home.Name)
				fmt.Println(game.Away.Name)
				fmt.Printf("Date: %s\n\n", date)
			}
		}
	}
}
