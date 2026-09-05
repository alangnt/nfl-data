package types

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
