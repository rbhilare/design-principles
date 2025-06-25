package main

import "fmt"

// Violating Interface Segregation Principle
type Operator interface {
	Reconcile() error
	Backup() error
	Failover() error
}

// DatabaseOperator implements everything
type DatabaseOperator struct{}

func (d DatabaseOperator) Reconcile() error {
	fmt.Println("Reconcile DB")
	return nil
}
func (d DatabaseOperator) Backup() error {
	fmt.Println("Backup DB")
	return nil
}
func (d DatabaseOperator) Failover() error {
	fmt.Println("Failover DB")
	return nil
}

// CacheOperator doesn’t need backup or failover, but must implement them
type CacheOperator struct{}

func (c CacheOperator) Reconcile() error {
	fmt.Println("Reconcile Cache")
	return nil
}
func (c CacheOperator) Backup() error {
	// Irrelevant, but required
	return nil
}
func (c CacheOperator) Failover() error {
	// Irrelevant, but required
	return nil
}

func Operators() {
	var op Operator = CacheOperator{}
	op.Reconcile()
}
