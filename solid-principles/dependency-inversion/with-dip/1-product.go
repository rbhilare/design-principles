package main

import "fmt"

// With Dependency Inversion Principle
// Define Product interface
type Product interface {
	Describe() string
}

type Ansible struct {
	ProductName string
}

func (a Ansible) Describe() string {
	return "Ansible Product Title: " + a.ProductName
}

type RHEL struct {
	ProductName string
}

func (r RHEL) Describe() string {
	return "RHEL Product Title: " + r.ProductName
}

// Function that depends on abstraction and accepts any Product
func PrintProduct(p Product) {
	fmt.Println(p.Describe())
}

func CallPrint() {
	ansible := Ansible{ProductName: "Ansible"}
	rhel := Ansible{ProductName: "RHEL"}

	// Inject different product implementations at runtime
	PrintProduct(ansible)
	PrintProduct(rhel)
}
