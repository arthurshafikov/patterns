package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/composite/solution/products"
)

func main() {
	box := NewBox([]products.HasPrice{
		products.NewHammer(),
		products.NewHammer(),
		products.NewTV(), // some new product
		NewBox([]products.HasPrice{
			products.NewNewspaper(),
			products.NewPhone(),
			NewBox([]products.HasPrice{
				NewBox([]products.HasPrice{
					products.NewNewspaper(),
					products.NewPhone(),
				}),
			}),
			products.NewPhone(),
			products.NewHammer(),
		}),
	})

	// we need to calculate the price.
	totalPrice := box.GetPrice()

	fmt.Printf("Total price is %v\n", totalPrice)
	// Total price is 112.
}
