package main

import "fmt"

func main() {
	// Example of using the make function to create a map of stock items
	stock := make(map[string]int)
	stock["item1"] = 10
	stock["item2"] = 5
	stock["item3"] = 20
	fmt.Println("stock:", stock)

	// Example of using the new keyword to create an instance of a product
	product := new(Product)
	product.Name = "T-shirt"
	product.Price = 29.99
	fmt.Println("product:", product)
}

// Definition of a custom type
type Product struct {
	Name  string
	Price float64
}
