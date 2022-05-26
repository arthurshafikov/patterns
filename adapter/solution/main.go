package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/adapter/solution/metrics"
	"github.com/arthurshafikov/patterns/adapter/solution/spaceship"
)

func main() {
	spaceXSpaceship := spaceship.NewSpaceship(28_968, 500)

	planets := []string{
		"Jupiter",
		"Mars",
		"Sun",
		"Moon",
		"Saturn",
	}

	for _, planet := range planets {
		var distance float64
		switch planet {
		case "Jupiter":
			distance = metrics.DistanceToTheJupiter()
		case "Mars":
			distance = metrics.DistanceToTheMars()
		case "Sun":
			distance = metrics.DistanceToTheSun()
		case "Moon":
			distance = metrics.DistanceToTheMoon()
		case "Saturn":
			distance = metrics.DistanceToTheSaturn()
		}
		estimateFlyHours := spaceXSpaceship.EstimateFlyHours(distance)
		estimateFuelNeeded := spaceXSpaceship.EstimateFuelAmount(distance)

		fmt.Printf("Estimate time to %s is %v hours, fuel needed: %vL\n", planet, estimateFlyHours, estimateFuelNeeded)
		/*
			Estimate time to Jupiter is 27824 hours, fuel needed: 1.3912e+07L
			Estimate time to Mars is 7779 hours, fuel needed: 3.8895e+06L
			Estimate time to Sun is 5167 hours, fuel needed: 2.5835e+06L
			Estimate time to Moon is 13 hours, fuel needed: 6500L
			Estimate time to Saturn is 41445 hours, fuel needed: 2.07225e+07L
		*/
	}
}
