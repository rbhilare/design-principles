package main

import "fmt"

// Violating/Without Dependency Inversion Principle
type Ansible struct {
	ProductName string
}

func (a Ansible) Describe() string {
	return "Ansible Product Title: " + a.ProductName
}

func PrintAnsibleProduct() {
	ansible := Ansible{ProductName: "Ansible"}
	fmt.Println(ansible.Describe())
}

type RHEL struct {
	ProductName string
}

func (r RHEL) Describe() string {
	return "RHEL Product Title: " + r.ProductName
}

func PrintRHELProduct() {
	rhel := Ansible{ProductName: "RHEL"}
	fmt.Println(rhel.Describe())
}

func CallPrint() {
	PrintAnsibleProduct() // Always works with Ansible, no flexibility
	PrintRHELProduct()    // Always works with RHEL, no flexibility
}
