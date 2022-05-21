package builders

import (
	"github.com/arthurshafikov/patterns/builder/solution/estates"
)

type HostelBuilder struct {
	hostel *estates.Hostel
}

func NewHostelBuilder() *HostelBuilder {
	return &HostelBuilder{
		hostel: estates.NewHostel(),
	}
}

func (h *HostelBuilder) SelectStyle(style string) EstateBuilder {
	h.hostel.Style = style

	return h
}

func (h *HostelBuilder) BuildRooms(numOfRooms int) EstateBuilder {
	h.hostel.NumOfRooms = numOfRooms

	return h
}

func (h *HostelBuilder) BuildSwimmingPool() EstateBuilder {
	h.hostel.HasSwimmingPool = true

	return h
}

func (h *HostelBuilder) BuildGarage() EstateBuilder {
	h.hostel.HasGarage = true

	return h
}

func (h *HostelBuilder) BuildGarden() EstateBuilder {
	h.hostel.HasGarden = true

	return h
}

func (h *HostelBuilder) BuildFancyStatues() EstateBuilder {
	h.hostel.HasFancyStatues = true

	return h
}

func (h *HostelBuilder) GetResult() estates.Estate {
	return h.hostel
}
