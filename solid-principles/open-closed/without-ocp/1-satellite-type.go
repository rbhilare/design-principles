package main

import "fmt"

type Satellite struct{}

// Violating Open Closed Principle
func (s Satellite) TransmitData(sType string) {
	if sType == "weather" {
		fmt.Println("Weather data transmission in progress.....")
	} else if sType == "gps" {
		fmt.Println("GPS data transmission in progress.....")
	} else if sType == "communication" {
		fmt.Println("Communication signals transmission in progress.....")
	} else {
		fmt.Println("Satellite type not know.....")
	}
}

func main() {
	satellite := Satellite{}
	satellite.TransmitData("weather")
	satellite.TransmitData("gps")
}
