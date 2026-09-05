package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
	Home    GameTeam `json:"home"`
	Away    GameTeam `json:"away"`
	Scoring Scoring  `json:"scoring"`
}

type Week struct {
	Games []Game `json:"games"`
}

type Schedule struct {
	Weeks []Week `json:"weeks"`
}

func GetSchedule(year *string, teamId *string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sportradar := os.Getenv("SPORTRADAR_API_KEY")
	if sportradar == "" {
		log.Fatal("API key required")
	}

	url := "https://api.sportradar.com/nfl/official/trial/v7/en/games/" + *year + "/REG/schedule.json"

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

	fmt.Println(result)
}
