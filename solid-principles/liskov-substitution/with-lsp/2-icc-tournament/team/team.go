package team

import (
	"fmt"

	"icc-tournament/player"
)

// Team struct
type Team struct {
	Name    string
	Players []player.Player
}

func (t Team) PlayMatch() []string {
	var performances []string
	for _, p := range t.Players {
		performances = append(performances, fmt.Sprintf("[%s] %s", t.Name, p.Play()))
	}
	return performances
}
