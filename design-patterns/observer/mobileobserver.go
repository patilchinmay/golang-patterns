package main

import "log"

//  mobileObserver implements observer interface
type mobileObserver struct {
	name string
}

func (m *mobileObserver) update(updatedTemperature int) {
	log.Println(m.name, " : updated temperature is ", updatedTemperature, " degree Celcius")
}
