package factories

type Factory interface {
	CreateChair() Chair
	CreateSofa() Sofa
	CreateCoffeeTable() CoffeeTable
}
