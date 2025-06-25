package main

import "fmt"

// Small, focused and more specific interfaces
// Adhering to Interface Segregation Principle
type Reconcile interface {
	Reconcile() error
}

type Backup interface {
	Backup() error
}

type Failover interface {
	Failover() error
}

// DatabaseOperator supports all features
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

// CacheOperator only supports reconciliation
type CacheOperator struct{}

func (c CacheOperator) Reconcile() error {
	fmt.Println("Reconcile Cache")
	return nil
}

// MetricsOperator collects metrics — no interfaces implemented here
type MetricsOperator struct{}

func (m MetricsOperator) CollectMetrics() {
	fmt.Println("Collecting metrics")
}

// Functions that operate only on relevant behavior
func DoReconcile(r Reconcile) {
	r.Reconcile()
}

func DoBackup(b Backup) {
	b.Backup()
}

func Opeators() {
	db := DatabaseOperator{}
	cache := CacheOperator{}

	DoReconcile(db)
	DoReconcile(cache)

	DoBackup(db)
	// DoBackup(cache) // CacheOperator doesn’t implement Backup
}
