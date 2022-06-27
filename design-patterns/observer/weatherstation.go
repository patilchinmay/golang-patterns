package main

//  weatherStation implements subject interface
type weatherStation struct {
	ObserverList []observer
}

func (w *weatherStation) registerObserver(o observer) {
	w.ObserverList = append(w.ObserverList, o)
}

func (w *weatherStation) notifyAll(updatedTemperature int) {
	for _, o := range w.ObserverList {
		o.update(updatedTemperature)
	}
}
