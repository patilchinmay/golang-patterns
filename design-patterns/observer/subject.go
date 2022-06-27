package main

type subject interface {
	registerObserver(o observer)
	// deregisterObserver()
	notifyAll(updatedTemperature int)
}
