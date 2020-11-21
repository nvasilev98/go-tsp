package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/nvasilev98/go-tsp/base"

	ga "github.com/nvasilev98/go-tsp/geneticAlgorithm"
)

var (

	// Define seed for default rand
	//seed = int64(1504372704)
	seed = time.Now().Unix()

	// Number of generation to loop through
	noGen = 100
	// Population Size
	popSize = 200
)

func main() {
	fmt.Println("Traveling sales person")
	//?
	log.SetOutput(ioutil.Discard)

	// Set seed value for default source
	rand.Seed(seed)
	fmt.Println("seed: ", seed)

	// Init TourManager
	tm := base.TourManager{}
	tm.NewTourManager()

	// Generate Cities
	var cities []base.City
	cities = *initRandomCities(20)

	// Add cities to TourManager
	for _, v := range cities {
		tm.AddCity(v)
	}

	tspGA(&tm, noGen)
}

// tspGA : Travelling sales person with genetic algorithm
// input :- TourManager, Number of generations
func tspGA(tm *base.TourManager, gen int) {
	p := base.Population{}
	p.InitPopulation(popSize, *tm)

	// Get initial fittest tour and it's tour distance
	fmt.Println("Start....")
	iFit := p.GetFittest()
	iTourDistance := iFit.TourDistance()
	//fmt.Println("Initial tour distance: ", iTourDistance)

	// Map to store fittest tours
	fittestTours := make([]base.Tour, 0, gen+1)
	fittestTours = append(fittestTours, *iFit)
	// Evolve population "gen" number of times
	for i := 1; i < gen+1; i++ {
		log.Println("Generation ", i)
		p = ga.EvolvePopulation(p)
		// Store fittest for each generation
		fittestTours = append(fittestTours, *p.GetFittest())
	}
	// Get final fittest tour and tour distance
	fFit := p.GetFittest()
	fTourDistance := fFit.TourDistance()

	// Store current best distance
	lastBestTourDistance := iTourDistance
	// Plot Generation 0

	for gn, t := range fittestTours {
		if t.TourDistance() < lastBestTourDistance {
			lastBestTourDistance = t.TourDistance()
		}
		if gn == 10 || gn == 13 || gn == 20 || gn == 30 {
			fmt.Printf("Generation %v: %v\n", gn, lastBestTourDistance)
		}
	}

	fmt.Println("Initial tour distance: ", iTourDistance)
	fmt.Println("Final tour distance: ", fTourDistance)
}

func initRandomCities(cityCount int) *[]base.City {
	cities := make([]base.City, 0, cityCount)

	for i := 0; i < cityCount; i++ {
		cities = append(cities, base.GenerateRandomCity())
	}
	return &cities
}
