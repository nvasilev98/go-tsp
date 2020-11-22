package base

import (
	"fmt"
	"math"
	"math/rand"
)

type City struct {
	x int
	y int
}

func GenerateRandomCity() City {
	c := City{}
	c.x = rand.Intn(100) * 10
	c.y = rand.Intn(100) * 10
	return c
}

func GenerateCity(x int, y int) City {
	c := City{}
	c.x = x
	c.y = y
	return c
}

func (a *City) SetLocation(x int, y int) {
	a.x = x
	a.y = y
}

func (a *City) DistanceTo(c City) float64 {
	idx := a.x - c.x
	idy := a.y - c.y

	if idx < 0 {
		idx = -idx
	}
	if idy < 0 {
		idy = -idy
	}

	fdx := float64(idx)
	fdy := float64(idy)

	fd := math.Sqrt((fdx * fdx) + (fdy * fdy))
	return fd
}

func (a *City) X() int {
	return a.x
}

func (a *City) Y() int {
	return a.y
}

func (a City) String() string {
	return fmt.Sprintf("{x%d y%d}", a.x, a.y)
}

func ShuffleCities(in []City) []City {
	out := make([]City, len(in), cap(in))
	perm := rand.Perm(len(in))
	for i, v := range perm {
		out[v] = in[i]
	}
	return out
}
