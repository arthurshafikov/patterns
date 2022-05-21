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

type AbstractEstate struct {
	Estate

	Style      string
	NumOfRooms int

	HasSwimmingPool bool
	HasGarage       bool
	HasGarden       bool
	HasFancyStatues bool
}

func AbstractNewEstate() *AbstractEstate {
	return &AbstractEstate{}
}

func (h AbstractEstate) PrintInfo() {
	text := fmt.Sprintf("Estate is built in %s style, has %v rooms", h.Style, h.NumOfRooms)

	var additionalInfo []string

	if h.HasSwimmingPool {
		additionalInfo = append(additionalInfo, "swimming pool")
	}
	if h.HasGarage {
		additionalInfo = append(additionalInfo, "garage")
	}
	if h.HasGarden {
		additionalInfo = append(additionalInfo, "garden")
	}
	if h.HasFancyStatues {
		additionalInfo = append(additionalInfo, "fancy statues")
	}

	if len(additionalInfo) > 0 {
		text += fmt.Sprintf(" and also has %s", strings.Join(additionalInfo, ", "))
	}

	fmt.Println(text)
}
