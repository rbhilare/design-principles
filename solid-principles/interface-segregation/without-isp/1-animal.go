package main

import "fmt"

// Violating Interface Segregation Principle
type Animal interface {
	Eat()
	Fly()
	Swim()
}

type Dog struct{}

func (d Dog) Eat() {
	fmt.Println("Dog can eat")
}

func (d Dog) Fly() {
	// Dogs can't fly
	panic("Dogs can't fly!")
}

func (d Dog) Swim() {
	fmt.Println("Dog can swim")
}

func main() {
	var a Animal = Dog{}
	a.Eat()
	a.Swim()
	a.Fly() // Will result in panic at runtime!
}
