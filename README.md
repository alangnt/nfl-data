# NFL Data in Go

This project in Go allows anyone to get the full active **roster** and **schedule** of any team in the NFL.

## Run the project

```go
go run .
```

## Go Linter

```bash
golangci-lint run
```

## Environment file

Must contain:

- SPORTRADAR_API_KEY

Get yours, it's free: [marketplace.sportradar.com](https://marketplace.sportradar.com/)

## Resources

- [developer.sportradar.com/football/reference/nfl-overview](https://developer.sportradar.com/football/reference/nfl-overview) - Spotradar's NFL data docs
- [github.com/charmbracelet/huh](https://github.com/charmbracelet/huh) - simple Go library for building interactive forms

## Project structure

```text
nfl-data/
├── README.md
├── go.mod
├── go.sum
├── hello.go
├── .env
├── types/
│   └── types.go
├── cards/
│   ├── game_card.go
│   └── player_card.go
└── utils/
    ├── env_utils.go
    ├── roster_utils.go
    ├── schedule_utils.go
    └── select_utils.go
```
