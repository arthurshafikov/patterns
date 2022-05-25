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
		var estimateFlyHours, estimateFuelNeeded float64
		switch planet {
		case "Jupiter":
			estimateFlyHours = spaceXSpaceship.EstimateFlyHours(metrics.DistanceToTheJupiter())
			estimateFuelNeeded = spaceXSpaceship.EstimateFuelAmount(metrics.DistanceToTheJupiter())
		case "Mars":
			estimateFlyHours = spaceXSpaceship.EstimateFlyHours(metrics.DistanceToTheMars())
			estimateFuelNeeded = spaceXSpaceship.EstimateFuelAmount(metrics.DistanceToTheMars())
		case "Sun":
			km := us_metrics.DistanceToTheSun() * kmInMile
			estimateFlyHours = spaceXSpaceship.EstimateFlyHours(km)
			estimateFuelNeeded = spaceXSpaceship.EstimateFuelAmount(km)
		case "Moon":
			km := us_metrics.DistanceToTheMoon() * kmInMile
			estimateFlyHours = spaceXSpaceship.EstimateFlyHours(km)
			estimateFuelNeeded = spaceXSpaceship.EstimateFuelAmount(km)
		case "Saturn":
			// A new programmer came and didn't know about miles convertion and wrote this code:
			km := us_metrics.DistanceToTheMoon()
			estimateFlyHours = spaceXSpaceship.EstimateFlyHours(km)
			estimateFuelNeeded = spaceXSpaceship.EstimateFuelAmount(km)
			// That's it, the ship would not have enough fuel for this journey and it'd crash.
		}

		fmt.Printf("Estimate time to %s is %v hours, fuel needed: %vL\n", planet, estimateFlyHours, estimateFuelNeeded)
	}
}
