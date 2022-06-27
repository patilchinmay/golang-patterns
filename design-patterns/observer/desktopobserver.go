package main

import "log"

//  desktopObserver implements observer interface
type desktopObserver struct {
	name string
}

func (d *desktopObserver) update(updatedTemperature int) {
	log.Println(d.name, " : updated temperature is ", updatedTemperature, " degree Celcius")
}
