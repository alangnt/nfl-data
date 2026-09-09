package main

import (
	"nfl-data/utils"
	"os"
)

func main() {
	sportradarKey, err := utils.GetSportradarAPIKey()
	if err != nil {
		os.Exit(1)
	}

	for {
		conference, err := utils.SelectConference()
		if err != nil {
			os.Exit(1)
		}

		if conference == "exit" {
			break
		}

		for {
			var teamID string

			if conference == "afc" {
				teamID, err = utils.SelectAFCTeamID()
				if err != nil {
					os.Exit(1)
				}
			} else {
				teamID, err = utils.SelectNFCTeamID()
				if err != nil {
					os.Exit(1)
				}
			}

			if teamID == "exit" {
				break
			}

			choice, err := utils.SelectTeamInfoChoice()
			if err != nil {
				os.Exit(1)
			}

			if choice == "roster" {
				err := utils.GetTeam(teamID, sportradarKey)
				if err != nil {
					os.Exit(1)
				}
			} else {
				for {
					year, err := utils.SelectYear()
					if err != nil {
						os.Exit(1)
					}

					if year == "exit" {
						break
					}

					for {
						seasonType, err := utils.SelectSeasonType()
						if err != nil {
							os.Exit(1)
						}

						if seasonType == "exit" {
							break
						}

						err = utils.GetSchedule(teamID, year, seasonType, sportradarKey)
						if err != nil {
							os.Exit(1)
						}
					}
				}
			}
		}
	}
}
