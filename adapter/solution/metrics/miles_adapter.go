package metrics

import "github.com/arthurshafikov/patterns/adapter/solution/us_metrics"

const kmInMile = 1.60934

// We adapted all the library functions to work as we need.
func DistanceToTheSun() float64 {
	return us_metrics.DistanceToTheSun() * kmInMile
}

func DistanceToTheMoon() float64 {
	return us_metrics.DistanceToTheMoon() * kmInMile
}

func DistanceToTheSaturn() float64 {
	return us_metrics.DistanceToTheSaturn() * kmInMile
}
