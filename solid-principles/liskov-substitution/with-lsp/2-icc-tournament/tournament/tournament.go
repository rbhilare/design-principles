package tournament

import (
	"fmt"

	"icc-tournament/team"
)

// Tournament struct
type Tournament struct {
	Name  string
	Teams []team.Team
}

func (t Tournament) Start() {
	fmt.Printf("ICC Tournament: %s\n\n", t.Name)

	for _, team := range t.Teams {
		fmt.Printf("Team: %s is playing...\n", team.Name)
		results := team.PlayMatch()
		for _, r := range results {
			fmt.Println("  " + r)
		}
		fmt.Println()
	}

	fmt.Println("Tournament Finished!")
}
