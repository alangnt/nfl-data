package cards

import (
	"fmt"
	"nfl-data/types"
	"time"
)

func DisplayGameCard(week *string, game *types.Game) {
	parsedTime, err := time.Parse(time.RFC3339, game.Scheduled)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	timeFormat := "Monday, Jan 2, 2006 at 3:04 PM"
	date := parsedTime.Format(timeFormat)

	fmt.Printf("Week %s\n", *week)
	fmt.Println(game.Home.Name)
	fmt.Println(&game.Away.Name)
	fmt.Printf("Date: %s\n\n", date)
}

func DisplayGameCards(games *map[string]types.Game) {
	for week, game := range *games {
		DisplayGameCard(&week, &game)
	}
}
