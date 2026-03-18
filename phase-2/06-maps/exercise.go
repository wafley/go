package main

import "fmt"

func main() {
	// Initial Data
	shop := map[string]int{
		"book": 5000,
		"pen": 3000,
		"ruler": 2000,
	}

	// Add
	shop["bag"] = 50000

	// Update
	shop["pen"] = 3500

	// Delete
	delete(shop, "ruler")

	// Transaction Logic
	shopping := []string{"book", "bag", "spidol"}
	total := 0
	for _, item := range shopping {
		price, exists := shop[item]

		if exists {
			fmt.Printf("Buy %s for %d\n", item, price)
			total += price
		} else {
			fmt.Printf("Sorry, %s is not available\n", item)
		}
	}
	fmt.Printf("Total payment: %d\n", total)
}