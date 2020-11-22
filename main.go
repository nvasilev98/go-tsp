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
	seed = time.Now().Unix()

	// Number of generation to loop through
	noGen = 100
	// Population Size
	popSize = 300
)

func main() {
	log.SetOutput(ioutil.Discard)
	rand.Seed(seed)

	tm := base.TourManager{}
	tm.NewTourManager()

	var cities []base.City

	var n int
	fmt.Println("Enter cities: ")
	fmt.Scanln(&n)
	cities = *initRandomCities(n)

	for _, v := range cities {
		tm.AddCity(v)
	}

	tspGA(&tm, noGen)
}

func tspGA(tm *base.TourManager, gen int) {
	p := base.Population{}
	p.InitPopulation(popSize, *tm)

	fmt.Println("Start....")
	iFit := p.GetFittest()
	iTourDistance := iFit.TourDistance()

	// Map to store fittest tours
	fittestTours := make([]base.Tour, 0, gen+1)
	fittestTours = append(fittestTours, *iFit)
	// Evolve population "gen" number of times
	for i := 1; i < gen+1; i++ {
		log.Println("Generation ", i)
		p = ga.EvolvePopulation(p)

		fittestTours = append(fittestTours, *p.GetFittest())
	}

	fFit := p.GetFittest()
	fTourDistance := fFit.TourDistance()

	lastBestTourDistance := iTourDistance

	for gn, t := range fittestTours {
		if t.TourDistance() < lastBestTourDistance {
			lastBestTourDistance = t.TourDistance()
		}
		if gn == 10 || gn == 25 || gn == 35 || gn == 50 {
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
