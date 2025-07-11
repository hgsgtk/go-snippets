package main

import (
	"fmt"

	"chef"
)

func main() {
	list := []chef.Dish{
		{Name: "Bruschetta", Type: "entree", PrepTime: 10},
		{Name: "Caesar Salad", Type: "entree", PrepTime: 8},
		{Name: "Steak", Type: "main", PrepTime: 25},
		{Name: "Pasta", Type: "main", PrepTime: 20},
		{Name: "Tiramisu", Type: "dessert", PrepTime: 12},
		{Name: "Gelato", Type: "dessert", PrepTime: 7},
	}
	  
	c := chef.NewChef(list)
	output := c.GetBestMeal()

	fmt.Printf("Best Meal:\n")
	fmt.Printf("  Entree: %s (%d minutes)\n", output.Entree.Name, output.Entree.PrepTime)
	fmt.Printf("  Main: %s (%d minutes)\n", output.Main.Name, output.Main.PrepTime)
	fmt.Printf("  Dessert: %s (%d minutes)\n", output.Dessert.Name, output.Dessert.PrepTime)
	fmt.Printf("  Total Time: %d minutes\n", output.TotalTime)
}