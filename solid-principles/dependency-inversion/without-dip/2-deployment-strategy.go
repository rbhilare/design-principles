package main

import "fmt"

// Violating/Without Dependency Inversion Principle
type Artifact struct {
	StrategyName string
}

func (a Artifact) Describe() string {
	return "Deployment strategy: " + a.StrategyName
}

func PrintArtifactStrategy() {
	artifact := Artifact{StrategyName: "Artifact"}
	fmt.Println(artifact.Describe())
}

type ReleaseVersion struct {
	StrategyName string
}

func (r ReleaseVersion) Describe() string {
	return "Deployment strategy: " + r.StrategyName
}

func PrintReleaseVersionStrategy() {
	releaseVersion := ReleaseVersion{StrategyName: "Release Version"}
	fmt.Println(releaseVersion.Describe())
}

func SelectDeploymentStrategy() {
	PrintArtifactStrategy()       // Always works with Artifact, no flexibility
	PrintReleaseVersionStrategy() // Always works with Release Version, no flexibility
}
