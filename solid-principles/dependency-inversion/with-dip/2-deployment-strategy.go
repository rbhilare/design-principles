package main

import "fmt"

// With Dependency Inversion Principle
// Define DeploymentStrategy interface
type DeploymentStrategy interface {
	Describe() string
}

type Artifact struct {
	StrategyName string
}

func (a Artifact) Describe() string {
	return "Deployment strategy: " + a.StrategyName
}

type ReleaseVersion struct {
	StrategyName string
}

func (r ReleaseVersion) Describe() string {
	return "Deployment strategy: " + r.StrategyName
}

// Function that depends on abstraction and accepts any DeploymentStrategy
func PrintDeploymentStrategy(d DeploymentStrategy) {
	fmt.Println(d.Describe())
}

func SelectDeploymentStrategy() {
	artifact := Artifact{StrategyName: "Artifact"}
	releaseVersion := ReleaseVersion{StrategyName: "Release Version"}

	// Inject different DeploymentStrategy implementations at runtime
	PrintDeploymentStrategy(artifact)
	PrintDeploymentStrategy(releaseVersion)
}
