package main

import (
	"icc-tournament/player"
	"icc-tournament/team"
	"icc-tournament/tournament"
)

func main() {
	// Players
	p1 := player.Batsman{PlayerName: "Rohit Sharma"}
	p2 := player.Bowler{PlayerName: "Jasprit Bumrah"}
	p3 := player.AllRounder{PlayerName: "Hardik Pandya"}

	p4 := player.Batsman{PlayerName: "Travis Head"}
	p5 := player.Bowler{PlayerName: "Mitchell Starc"}
	p6 := player.AllRounder{PlayerName: "Pat Cummins"}

	// Teams
	india := team.Team{
		Name:    "India",
		Players: []player.Player{p1, p2, p3},
	}

	australia := team.Team{
		Name:    "Australia",
		Players: []player.Player{p4, p5, p6},
	}

	// Tournament
	tournamentCT := tournament.Tournament{
		Name:  "ICC Champions Trophy 2025",
		Teams: []team.Team{india, australia},
	}

	// Start ICC Tournament
	tournamentCT.Start()
}
