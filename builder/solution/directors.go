package main

import (
	"github.com/arthurshafikov/patterns/builder/solution/builders"
	"github.com/arthurshafikov/patterns/builder/solution/estates"
)

const (
	CheapCost     = "cheap"
	NormalCost    = "normal"
	ExpensiveCost = "expensive"
)

type Director struct {
	builder builders.Builder
}

func NewDirector() *Director {
	return &Director{}
}

func (d *Director) SetBuilder(builder builders.Builder) *Director {
	d.builder = builder

	return d
}

func (d *Director) Create(cost string) estates.Estate {
	switch cost {
	case CheapCost:
		return d.builder.
			SelectStyle(estates.BungalowStyle).
			BuildRooms(2).
			GetResult()
	case NormalCost:
		return d.builder.
			SelectStyle(estates.ModernStyle).
			BuildRooms(5).
			BuildGarage().
			GetResult()
	case ExpensiveCost:
		return d.builder.
			SelectStyle(estates.CottageStyle).
			BuildRooms(15).
			BuildSwimmingPool().
			BuildGarage().
			BuildGarden().
			BuildFancyStatues().
			GetResult()
	default:
		panic("cost is not defined")
	}
}
