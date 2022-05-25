package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/adapter/problem/metrics"
	"github.com/arthurshafikov/patterns/adapter/problem/spaceship"
	"github.com/arthurshafikov/patterns/adapter/problem/us_metrics"
)

const kmInMile = 1.60934

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
			distance = us_metrics.DistanceToTheSun() * kmInMile
		case "Moon":
			distance = us_metrics.DistanceToTheMoon() * kmInMile
		case "Saturn":
			// A new programmer came and didn't know about miles convertion and wrote this code:
			distance = us_metrics.DistanceToTheSaturn()
			// That's it, the ship would not have enough fuel for this journey and it'd crash.
		}
		estimateFlyHours := spaceXSpaceship.EstimateFlyHours(distance)
		estimateFuelNeeded := spaceXSpaceship.EstimateFuelAmount(distance)

		fmt.Printf("Estimate time to %s is %v hours, fuel needed: %vL\n", planet, estimateFlyHours, estimateFuelNeeded)
		/*
			Estimate time to Jupiter is 27824 hours, fuel needed: 1.3912e+07L
			Estimate time to Mars is 7779 hours, fuel needed: 3.8895e+06L
			Estimate time to Sun is 5167 hours, fuel needed: 2.5835e+06L
			Estimate time to Moon is 13 hours, fuel needed: 6500L
			Estimate time to Saturn is 25753 hours, fuel needed: 1.28765e+07L
		*/
	}
}
