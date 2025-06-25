package main

import "fmt"

// Small, focused and more specific interfaces
// Adhering to Interface Segregation Principle
type IEat interface {
	Eat()
}

type IFly interface {
	Fly()
}

type ISwim interface {
	Swim()
}

// Dog implements only relevant interfaces i.e. IEat and ISwim
type Dog struct{}

func (d Dog) Eat() {
	fmt.Println("Dog can eat")
}

func (d Dog) Swim() {
	fmt.Println("Dog can swim")
}

// Bird implements only relevant interfaces i.e. IEat and IFly
type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("Bird can eat")
}

func (b Bird) Fly() {
	fmt.Println("Bird can fly")
}

// Fish implements only relevant interfaces i.e. IEat and ISwim
type Fish struct{}

func (f Fish) Eat() {
	fmt.Println("Fish can eat")
}

func (f Fish) Swim() {
	fmt.Println("Fish can swim")
}

func main() {
	var e IEat = Dog{}
	e.Eat()

	var s ISwim = Dog{}
	s.Swim()

	var f IFly = Bird{}
	f.Fly()
}
