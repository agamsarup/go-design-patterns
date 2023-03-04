package behavioural

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	// Define some sample books and magazines
	b1 := Book{Title: "Design Patterns", Author: "Erich Gamma"}
	b2 := Book{Title: "Clean Code", Author: "Robert Martin"}
	m1 := Magazine{Title: "Scientific American", Publisher: "Nature Publishing Group"}
	m2 := Magazine{Title: "The Economist", Publisher: "The Economist Group"}

	// Create an array of Visitable items
	items := []Visitable{b1, b2, m1, m2}

	// Create an instance of the InventoryVisitor struct
	iv := InventoryVisitor{}

	// Iterate over each item in the array and call the Accept method
	for _, item := range items {
		item.Accept(&iv)
	}

	// Print the inventory count
	fmt.Printf("Inventory count: %d books, %d magazines\n", iv.NumBooks, iv.NumMagazines)
}
