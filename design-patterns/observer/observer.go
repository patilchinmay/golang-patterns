package main

type observer interface {
	update(updatedTemperature int)
}
