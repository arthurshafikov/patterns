package factories

type Factory interface {
	CreateChair() Chair
	CreateSofa() Chair
	CreateCoffeeTable() Chair
}
