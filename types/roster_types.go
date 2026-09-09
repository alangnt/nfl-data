package types

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
