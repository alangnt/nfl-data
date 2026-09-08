package main

import (
	"nfl-data/utils"

	_ "github.com/lib/pq"
)

func main() {
	for {
		conference := utils.SelectConference()

		if conference == "exit" {
			break
		}

		for {
			var teamId string

			if conference == "afc" {
				teamId = utils.SelectAFCTeamId()
			} else {
				teamId = utils.SelectNFCTeamId()
			}

			if teamId == "exit" {
				break
			}

			choice := utils.SelectTeamInfoChoice()
			if choice == "roster" {
				utils.GetTeam(teamId)
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

						utils.GetSchedule(teamId, year, seasonType)
					}
				}
			}
		}
	}
}
