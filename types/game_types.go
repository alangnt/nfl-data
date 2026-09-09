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
	Title     string   `json:"title"`
}

type Week struct {
	Games    []Game `json:"games"`
	Sequence int    `json:"sequence"`
}

type Schedule struct {
	Weeks []Week `json:"weeks"`
}
