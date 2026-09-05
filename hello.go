package main

import (
	"nfl-data/utils"

	_ "github.com/lib/pq"
)

func main() {
	conference := utils.GetConference()

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
		utils.GetSchedule(&teamId)
	}
}
