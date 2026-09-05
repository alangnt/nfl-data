package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

func GetYear() string {
	var year string

	err := huh.NewSelect[string]().
		Title("What year?").
		Options(
			huh.NewOption("2026", "2026"),
			huh.NewOption("2025", "2025"),
		).
		Value(&year).
		Run()

	if err != nil {
		fmt.Println("Selection cancelled:", err)
		os.Exit(1)
	}

	return year
}

func GetConference() string {
	var conference string

	err := huh.NewSelect[string]().
		Title("Choose your conference").
		Options(
			huh.NewOption("AFC", "afc"),
			huh.NewOption("NFC", "nfc"),
		).
		Value(&conference).
		Run()

	if err != nil {
		fmt.Println("Selection cancelled:", err)
		os.Exit(1)
	}

	return conference
}

func GetAFCTeamId() string {
	var teamId string

	err := huh.NewSelect[string]().
		Title("Choose your team").
		Options(
			huh.NewOption("Baltimore Ravens", "ebd87119-b331-4469-9ea6-d51fe3ce2f1c"),
			huh.NewOption("Buffalo Bills", "768c92aa-75ff-4a43-bcc0-f2798c2e1724"),
			huh.NewOption("Cincinnati Bengals", "ad4ae08f-d808-42d5-a1e6-e9bc4e34d123"),
			huh.NewOption("Cleveland Browns", "d5a2eb42-8065-4174-ab79-0a6fa820e35e"),
			huh.NewOption("Denver Broncos", "ce92bd47-93d5-4fe9-ada4-0fc681e6caa0"),
			huh.NewOption("Houston Texans", "82d2d380-3834-4938-835f-aec541e5ece7"),
			huh.NewOption("Indianapolis Colts", "82cf9565-6eb9-4f01-bdbd-5aa0d472fcd9"),
			huh.NewOption("Jacksonville Jaguars", "f7ddd7fa-0bae-4f90-bc8e-669e4d6cf2de"),
			huh.NewOption("Kansas City Chiefs", "6680d28d-d4d2-49f6-aace-5292d3ec02c2"),
			huh.NewOption("Las Vegas Raiders", "7d4fcc64-9cb5-4d1b-8e75-8a906d1e1576"),
			huh.NewOption("Los Angeles Chargers", "1f6dcffb-9823-43cd-9ff4-e7a8466749b5"),
			huh.NewOption("Miami Dolphins", "4809ecb0-abd3-451d-9c4a-92a90b83ca06"),
			huh.NewOption("New England Patriots", "97354895-8c77-4fd4-a860-32e62ea7382a"),
			huh.NewOption("New York Jets", "5fee86ae-74ab-4bdd-8416-42a9dd9964f3"),
			huh.NewOption("Pittsburgh Steelers", "cb2f9f1f-ac67-424e-9e72-1475cb0ed398"),
			huh.NewOption("Tennessee Titans", "d26a1ca5-722d-4274-8f97-c92e49c96315"),
		).
		Value(&teamId).
		Run()

	if err != nil {
		fmt.Println("Selection cancelled:", err)
		os.Exit(1)
	}

	return teamId
}

func GetNFCTeamId() string {
	var teamId string

	err := huh.NewSelect[string]().
		Title("Choose your team").
		Options(
			huh.NewOption("Arizona Cardinals", "de760528-1dc0-416a-a978-b510d20692ff"),
			huh.NewOption("Atlanta Falcons", "e6aa13a4-0055-48a9-bc41-be28dc106929"),
			huh.NewOption("Carolina Panthers", "f14bf5cc-9a82-4a38-bc15-d39f75ed5314"),
			huh.NewOption("Chicago Bears", "7b112545-38e6-483c-a55c-96cf6ee49cb8"),
			huh.NewOption("Dallas Cowboys", "e627eec7-bbae-4fa4-8e73-8e1d6bc5c060"),
			huh.NewOption("Detroit Lions", "c5a59daa-53a7-4de0-851f-fb12be893e9e"),
			huh.NewOption("Green Bay Packers", "a20471b4-a8d9-40c7-95ad-90cc30e46932"),
			huh.NewOption("Los Angeles Rams", "2eff2a03-54d4-46ba-890e-2bc3925548f3"),
			huh.NewOption("Minnesota Vikings", "33405046-04ee-4058-a950-d606f8c30852"),
			huh.NewOption("New Orleans Saints", "0d855753-ea21-4953-89f9-0e20aff9eb73"),
			huh.NewOption("New York Giants", "04aa1c9d-66da-489d-b16a-1dee3f2eec4d"),
			huh.NewOption("Philadelphia Eagles", "386bdbf9-9eea-4869-bb9a-274b0bc66e80"),
			huh.NewOption("San Francisco 49ers", "f0e724b0-4cbf-495a-be47-013907608da9"),
			huh.NewOption("Seattle Seahawks", "3d08af9e-c767-4f88-a7dc-b920c6d2b4a8"),
			huh.NewOption("Tampa Bay Buccaneers", "4254d319-1bc7-4f81-b4ab-b5e6f3402b69"),
			huh.NewOption("Washington Commanders", "22052ff7-c065-42ee-bc8f-c4691c50e624"),
		).
		Value(&teamId).
		Run()

	if err != nil {
		fmt.Println("Selection cancelled:", err)
		os.Exit(1)
	}

	return teamId
}

func GetTeamInfoChoice() string {
	var choice string

	err := huh.NewSelect[string]().
		Title("What do you want to know?").
		Options(
			huh.NewOption("Roster", "roster"),
			huh.NewOption("Schedule", "schedule"),
		).
		Value(&choice).
		Run()

	if err != nil {
		fmt.Println("Selection cancelled:", err)
		os.Exit(1)
	}

	return choice
}
