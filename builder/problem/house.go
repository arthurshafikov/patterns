package main

import (
	"fmt"
	"strings"
)

const (
	ModernStyle   = "modern"
	CottageStyle  = "cottage"
	BungalowStyle = "bungalow"
)

type House struct {
	style      string
	numOfRooms int

	hasSwimmingPool bool
	hasGarage       bool
	hasGarden       bool
	hasFancyStatues bool
}

// Huge constructor with a lot of params that sometimes are redunant.
func NewHouse(
	style string,
	numOfRooms int,
	hasSwimmingPool,
	hasGarage,
	hasGarden,
	hasFancyStatues bool,
) *House {
	return &House{
		style:      style,
		numOfRooms: numOfRooms,

		hasSwimmingPool: hasSwimmingPool,
		hasGarage:       hasGarage,
		hasGarden:       hasGarden,
		hasFancyStatues: hasFancyStatues,
	}
}

func (h House) PrintInfo() {
	text := fmt.Sprintf("House is built in %s style, has %v rooms", h.style, h.numOfRooms)

	var additionalInfo []string

	if h.hasSwimmingPool {
		additionalInfo = append(additionalInfo, "swimming pool")
	}
	if h.hasGarage {
		additionalInfo = append(additionalInfo, "garage")
	}
	if h.hasGarden {
		additionalInfo = append(additionalInfo, "garden")
	}
	if h.hasFancyStatues {
		additionalInfo = append(additionalInfo, "fancy statues")
	}

	if len(additionalInfo) > 0 {
		text += fmt.Sprintf(" and also has %s", strings.Join(additionalInfo, ", "))
	}

	fmt.Println(text)
}
