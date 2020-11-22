package base

import (
	"strconv"
)

type Tour struct {
	tourCities []City
	fitness    float64
	distance   float64
}

func (a *Tour) InitTour(numberOfCities int) {
	a.tourCities = make([]City, numberOfCities)
}

func (a *Tour) InitTourCities(tm TourManager) {
	a.InitTour(tm.NumberOfCities())

	for i := 0; i < tm.NumberOfCities(); i++ {
		a.SetCity(i, tm.GetCity(i))
	}

	a.tourCities = ShuffleCities(a.tourCities)
}

func (a *Tour) GetCity(tourPosition int) City {
	return a.tourCities[tourPosition]
}

func (a *Tour) SetCity(tourPosition int, c City) {
	a.tourCities[tourPosition] = c

	a.fitness = 0
	a.distance = 0
}

func (a *Tour) ResetFitnessDistance() {
	a.fitness = 0
	a.distance = 0
}

func (a *Tour) TourSize() int {
	return len(a.tourCities)
}

func (a *Tour) TourDistance() float64 {
	if a.distance == 0 {
		td := float64(0)
		for i := 0; i < a.TourSize(); i++ {
			fromC := a.GetCity(i)
			var destC City
			if i+1 < a.TourSize() {
				destC = a.GetCity(i + 1)
			} else {
				destC = a.GetCity(0)
			}
			td += fromC.DistanceTo(destC)
		}
		a.distance = td
	}
	return a.distance
}

func (a *Tour) Fitness() float64 {
	if a.fitness == 0 {
		a.fitness = 1 / a.TourDistance()
	}
	return a.fitness
}

func (a *Tour) ContainCity(c City) bool {
	for _, cs := range a.tourCities {
		if cs == c {
			return true
		}
	}
	return false
}

func (a Tour) String() string {
	s := "|"
	for i, c := range a.tourCities {
		s += strconv.Itoa(i) + c.String() + "|"
	}
	return s
}
