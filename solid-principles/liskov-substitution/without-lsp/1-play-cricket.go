package main

import "fmt"

// Violating Liskov Substitution Principle

// We will use interfaces for players
type Player interface {
	PlayMatch() string
}

// Batsman implements Player
type Batsman struct{}

func (b Batsman) PlayMatch() string {
	return "Batsman is batting"
}

// Commentator also implements Player (which is wrong)
type Commentator struct{}

func (c Commentator) PlayMatch() string {
	panic("Commentator does not play in matches!")
}

// SimulateMatch function that uses the interface
func SimulateMatch(p Player) {
	fmt.Println(p.PlayMatch())
}

func main() {
	var p1 Player = Batsman{}
	var p2 Player = Commentator{}

	SimulateMatch(p1) // This will work
	SimulateMatch(p2) // This will results in Panics at runtime! violating LSP
}
