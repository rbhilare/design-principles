package main

import "fmt"

// With Liskov Substitution Principle

// We will use interfaces for players
type Player interface {
	PlayMatch() string
}

// We will use interfaces for commentators
type Commentator interface {
	Commentate() string
}

// Batsman implements Player
type Batsman struct{}

func (b Batsman) PlayMatch() string {
	return "Batsman is batting"
}

// Bowler implements Player
type Bowler struct{}

func (b Bowler) PlayMatch() string {
	return "Bowler is bowling"
}

// SimulateMatch function that uses the interface
func SimulateMatch(p Player) {
	fmt.Println(p.PlayMatch())
}

func main() {
	var p1 Player = Batsman{}
	var p2 Player = Bowler{}

	SimulateMatch(p1) // It Outputs: Batsman is batting
	SimulateMatch(p2) // It Outputs: Bowler is bowling
}
