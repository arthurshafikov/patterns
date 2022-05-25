package spaceship

import "math"

type Spaceship struct {
	avgSpeed        float64 // km / hours
	fuelConsPerHour float64 // L / hours
}

func NewSpaceship(avgSpeed, fuelConsPerHour float64) *Spaceship {
	return &Spaceship{
		avgSpeed:        avgSpeed,
		fuelConsPerHour: fuelConsPerHour,
	}
}

func (s *Spaceship) EstimateFlyHours(distanceInKm float64) float64 {
	return math.Round(distanceInKm / s.avgSpeed)
}

func (s *Spaceship) EstimateFuelAmount(distanceInKm float64) float64 {
	return s.fuelConsPerHour * s.EstimateFlyHours(distanceInKm)
}
