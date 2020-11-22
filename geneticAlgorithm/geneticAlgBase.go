package geneticAlgorithm

import (
	"math/rand"

	"github.com/nvasilev98/go-tsp/base"
)

//Parameters
var (
	mutationRate        float64 = 0.015
	tournamentSize      int     = 5
	elitism             bool    = true
	randomCrossoverRate         = false
	defCrossoverRate    float32 = 0.7
)

func CrossoverRate() float32 {
	if randomCrossoverRate {
		return rand.Float32()
	}
	return defCrossoverRate
}

func Crossover(p1 base.Tour, p2 base.Tour) base.Tour {
	size := p1.TourSize()
	// Child Tour
	c := base.Tour{}
	c.InitTour(size)

	nc := int(CrossoverRate() * float32(size))
	if nc == 0 {
		return p1
	}

	sp := int(rand.Float32() * float32(size))

	ep := (sp + nc) % size

	p2s := make([]int, 0, size-nc)

	if sp < ep {
		for i := 0; i < size; i++ {
			if i >= sp && i < ep {
				c.SetCity(i, p1.GetCity(i))
			} else {
				p2s = append(p2s, i)
			}
		}
	} else if sp > ep {
		for i := 0; i < size; i++ {
			if !(i >= ep && i < sp) {
				c.SetCity(i, p1.GetCity(i))
			} else {
				p2s = append(p2s, i)
			}
		}
	}

	j := 0

	for i := 0; i < size; i++ {

		if !c.ContainCity(p2.GetCity(i)) {
			c.SetCity(p2s[j], p2.GetCity(i))
			j++
		}
	}
	return c
}

// Chance of mutation for each City based on mutation rate
func Mutation(in *base.Tour) {

	for p1 := 0; p1 < in.TourSize(); p1++ {
		if rand.Float64() < mutationRate {

			p2 := int(float64(in.TourSize()) * rand.Float64())

			c1 := in.GetCity(p1)
			c2 := in.GetCity(p2)

			in.SetCity(p1, c2)
			in.SetCity(p2, c1)
		}
	}
}

// TournamentSelection : select a group at random and pick the best parent
func TournamentSelection(pop base.Population) base.Tour {
	tourny := base.Population{}
	tourny.InitEmpty(tournamentSize)

	for i := 0; i < tournamentSize; i++ {
		r := int(rand.Float64() * float64(pop.PopulationSize()))
		tourny.SaveTour(i, *pop.GetTour(r))
	}
	fTour := tourny.GetFittest()
	return *fTour
}

func EvolvePopulation(pop base.Population) base.Population {
	npop := base.Population{}
	npop.InitEmpty(pop.PopulationSize())

	popOffset := 0
	if elitism {
		npop.SaveTour(0, *pop.GetFittest())
		popOffset = 1
	}

	for i := popOffset; i < npop.PopulationSize(); i++ {
		p1 := TournamentSelection(pop)
		p2 := TournamentSelection(pop)
		child := Crossover(p1, p2)
		Mutation(&child)
		npop.SaveTour(i, child)
	}
	return npop
}
