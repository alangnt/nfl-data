package main

import (
	"nfl-data/utils"

	_ "github.com/lib/pq"
)

func main() {
	for {
		conference := utils.GetConference()

		if conference == "exit" {
			break
		}

		var teamId string

		if conference == "afc" {
			teamId = utils.GetAFCTeamId()
		} else {
			teamId = utils.GetNFCTeamId()
		}

		choice := utils.GetTeamInfoChoice()
		if choice == "roster" {
			utils.GetTeam(&teamId)
		} else {
			for {
				year := utils.GetYear()

				if year == "exit" {
					break
				}

				utils.GetSchedule(&teamId, &year)
			}
		}
	}
}
