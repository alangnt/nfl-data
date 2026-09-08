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

		for {
			var teamId string

			if conference == "afc" {
				teamId = utils.GetAFCTeamId()
			} else {
				teamId = utils.GetNFCTeamId()
			}

			if teamId == "exit" {
				break
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

					for {
						seasonType := utils.GetSeasonType()
						if seasonType == "exit" {
							break
						}

						utils.GetSchedule(&teamId, &year, &seasonType)
					}
				}
			}
		}
	}
}
