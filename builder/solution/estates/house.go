package estates

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

func NewHouse() *House {
	return &House{}
}

func (h *House) SelectStyle(style string) *House {
	h.style = style

	return h
}

func (h *House) BuildRooms(numOfRooms int) *House {
	h.numOfRooms = numOfRooms

	return h
}

func (h *House) BuildSwimmingPool() *House {
	h.hasSwimmingPool = true

	return h
}

func (h *House) BuildGarage() *House {
	h.hasGarage = true

	return h
}

func (h *House) BuildGarden() *House {
	h.hasGarden = true

	return h
}

func (h *House) BuildFancyStatues() *House {
	h.hasFancyStatues = true

	return h
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
