package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	MobileObserver := &mobileObserver{name: "mobileObserver"}
	DesktopObserver := &desktopObserver{name: "desktopObserver"}

	WeatherStation := weatherStation{}

	WeatherStation.registerObserver(MobileObserver)
	WeatherStation.registerObserver(DesktopObserver)

	updateTemperature(WeatherStation)
}

// Update random temperature every 1 second
// Stop after 5 seconds
func updateTemperature(w weatherStation) {
	updateTimeTicker := time.NewTicker(1 * time.Second)
	defer func() { updateTimeTicker.Stop() }()

	timer := time.After(5 * time.Second)

	for {
		select {
		case <-updateTimeTicker.C:
			updatedTemperature := getRandomTemperature()
			w.notifyAll(updatedTemperature)
		case <-timer:
			fmt.Println("5 seconds! End Ticker")
			return
		}
	}
}

func getRandomTemperature() int {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 30
	updatedTemperature := rand.Intn(max-min+1) + min

	return updatedTemperature
}
