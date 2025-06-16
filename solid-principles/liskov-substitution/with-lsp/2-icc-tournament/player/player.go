package player

import "fmt"

// Player interface
type Player interface {
	Name() string
	Play() string
}

// Batsman struct
type Batsman struct {
	PlayerName string
}

func (b Batsman) Name() string {
	return b.PlayerName
}

func (b Batsman) Play() string {
	return fmt.Sprintf("%s scores 50 runs", b.PlayerName)
}

// Bowler struct
type Bowler struct {
	PlayerName string
}

func (b Bowler) Name() string {
	return b.PlayerName
}

func (b Bowler) Play() string {
	return fmt.Sprintf("%s takes 4 wickets", b.PlayerName)
}

// AllRounder struct
type AllRounder struct {
	PlayerName string
}

func (a AllRounder) Name() string {
	return a.PlayerName
}

func (a AllRounder) Play() string {
	return fmt.Sprintf("%s scores 90 runs and takes 2 wickets", a.PlayerName)
}
