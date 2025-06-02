package main

import "fmt"

// With Open Closed Principle

// We will use interfaces for each satellite
type Satellite interface {
	Transmit()
}

type WeatherSatellite struct{}

func (ws WeatherSatellite) Transmit() {
	fmt.Println("Weather data transmission in progress.....")

}

type GPSSatellite struct{}

func (gs GPSSatellite) Transmit() {
	fmt.Println("GPS data transmission in progress.....")

}

type CommunicationSatellite struct{}

func (cs CommunicationSatellite) Transmit() {
	fmt.Println("Communication signals transmission in progress.....")

}

type GroundStation struct{}

func (gs GroundStation) Receive(s Satellite) {
	s.Transmit()
}

func main() {

	gs := GroundStation{}

	weather := WeatherSatellite{}
	gps := GPSSatellite{}
	communication := CommunicationSatellite{}

	gs.Receive(weather)
	gs.Receive(gps)
	gs.Receive(communication)

}
