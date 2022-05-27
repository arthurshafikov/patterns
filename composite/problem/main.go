package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/composite/problem/products"
)

func main() {
	box := NewBox([]any{
		products.NewHammer(),
		products.NewHammer(),
		NewBox([]any{
			products.NewNewspaper(),
			products.NewPhone(),
			NewBox([]any{
				NewBox([]any{
					products.NewNewspaper(),
					products.NewPhone(),
				}),
			}),
			products.NewPhone(),
			products.NewHammer(),
		}),
	})

	// we need to calculate the price.
	allProducts := box.OpenBox()
	totalPrice := 0

	for _, product := range allProducts {
		if hammer, ok := product.(*products.Hammer); ok {
			totalPrice += hammer.GetPrice()
		}
		if newspaper, ok := product.(*products.Newspaper); ok {
			totalPrice += newspaper.GetPrice()
		}
		if phone, ok := product.(*products.Phone); ok {
			totalPrice += phone.GetPrice()
		}
	}

	fmt.Printf("Total price is %v\n", totalPrice)
	// Total price is 62.
}
