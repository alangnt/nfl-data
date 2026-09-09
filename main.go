package main

import (
	"nfl-data/utils"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	sportradarKey, err := utils.GetSportradarAPIKey()
	if err != nil {
		os.Exit(1)
	}

	for {
		conference := utils.SelectConference()

		if conference == "exit" {
			break
		}

		for {
			var teamID string

			if conference == "afc" {
				teamID = utils.SelectAFCTeamID()
			} else {
				teamID = utils.SelectNFCTeamID()
			}

			if teamID == "exit" {
				break
			}

			choice := utils.SelectTeamInfoChoice()
			if choice == "roster" {
				utils.GetTeam(teamID, sportradarKey)
			} else {
				for {
					year := utils.SelectYear()
					if year == "exit" {
						break
					}

					for {
						seasonType := utils.SelectSeasonType()
						if seasonType == "exit" {
							break
						}

						utils.GetSchedule(teamID, year, seasonType, sportradarKey)
					}
				}
			}
		}
	}
}
