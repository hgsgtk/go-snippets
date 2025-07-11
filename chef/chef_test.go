package chef_test

import (
	"testing"

	"chef"
)

func TestGetBestMeal(t *testing.T) {
	tests := []struct {
		name     string
		dishes   []chef.Dish
		expected chef.ChefRecipe
	}{
		{
			name: "Basic case - selects minimum prep time dishes",
			dishes: []chef.Dish{
				{Name: "Bruschetta", Type: "entree", PrepTime: 10},
				{Name: "Caesar Salad", Type: "entree", PrepTime: 8},
				{Name: "Steak", Type: "main", PrepTime: 25},
				{Name: "Pasta", Type: "main", PrepTime: 20},
				{Name: "Tiramisu", Type: "dessert", PrepTime: 12},
				{Name: "Gelato", Type: "dessert", PrepTime: 7},
			},
			expected: chef.ChefRecipe{
				Entree:    chef.Dish{Name: "Caesar Salad", Type: "entree", PrepTime: 8},
				Main:      chef.Dish{Name: "Pasta", Type: "main", PrepTime: 20},
				Dessert:   chef.Dish{Name: "Gelato", Type: "dessert", PrepTime: 7},
				TotalTime: 35,
			},
		},
		{
			name: "Single dish per type",
			dishes: []chef.Dish{
				{Name: "Soup", Type: "entree", PrepTime: 15},
				{Name: "Chicken", Type: "main", PrepTime: 30},
				{Name: "Cake", Type: "dessert", PrepTime: 20},
			},
			expected: chef.ChefRecipe{
				Entree:    chef.Dish{Name: "Soup", Type: "entree", PrepTime: 15},
				Main:      chef.Dish{Name: "Chicken", Type: "main", PrepTime: 30},
				Dessert:   chef.Dish{Name: "Cake", Type: "dessert", PrepTime: 20},
				TotalTime: 65,
			},
		},
		{
			name: "Multiple dishes with same prep time - selects first encountered",
			dishes: []chef.Dish{
				{Name: "Salad A", Type: "entree", PrepTime: 10},
				{Name: "Salad B", Type: "entree", PrepTime: 10},
				{Name: "Fish A", Type: "main", PrepTime: 20},
				{Name: "Fish B", Type: "main", PrepTime: 20},
				{Name: "Ice Cream A", Type: "dessert", PrepTime: 5},
				{Name: "Ice Cream B", Type: "dessert", PrepTime: 5},
			},
			expected: chef.ChefRecipe{
				Entree:    chef.Dish{Name: "Salad A", Type: "entree", PrepTime: 10},
				Main:      chef.Dish{Name: "Fish A", Type: "main", PrepTime: 20},
				Dessert:   chef.Dish{Name: "Ice Cream A", Type: "dessert", PrepTime: 5},
				TotalTime: 35,
			},
		},
		{
			name: "All dishes have same prep time",
			dishes: []chef.Dish{
				{Name: "Appetizer", Type: "entree", PrepTime: 15},
				{Name: "Dinner", Type: "main", PrepTime: 15},
				{Name: "Sweet", Type: "dessert", PrepTime: 15},
			},
			expected: chef.ChefRecipe{
				Entree:    chef.Dish{Name: "Appetizer", Type: "entree", PrepTime: 15},
				Main:      chef.Dish{Name: "Dinner", Type: "main", PrepTime: 15},
				Dessert:   chef.Dish{Name: "Sweet", Type: "dessert", PrepTime: 15},
				TotalTime: 45,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := chef.NewChef(tt.dishes)
			result := c.GetBestMeal()
			
			// Check entree
			if result.Entree.Name != tt.expected.Entree.Name {
				t.Errorf("Entree name: got %s, want %s", result.Entree.Name, tt.expected.Entree.Name)
			}
			if result.Entree.PrepTime != tt.expected.Entree.PrepTime {
				t.Errorf("Entree prep time: got %d, want %d", result.Entree.PrepTime, tt.expected.Entree.PrepTime)
			}
			
			// Check main
			if result.Main.Name != tt.expected.Main.Name {
				t.Errorf("Main name: got %s, want %s", result.Main.Name, tt.expected.Main.Name)
			}
			if result.Main.PrepTime != tt.expected.Main.PrepTime {
				t.Errorf("Main prep time: got %d, want %d", result.Main.PrepTime, tt.expected.Main.PrepTime)
			}
			
			// Check dessert
			if result.Dessert.Name != tt.expected.Dessert.Name {
				t.Errorf("Dessert name: got %s, want %s", result.Dessert.Name, tt.expected.Dessert.Name)
			}
			if result.Dessert.PrepTime != tt.expected.Dessert.PrepTime {
				t.Errorf("Dessert prep time: got %d, want %d", result.Dessert.PrepTime, tt.expected.Dessert.PrepTime)
			}
			
			// Check total time
			if result.TotalTime != tt.expected.TotalTime {
				t.Errorf("Total time: got %d, want %d", result.TotalTime, tt.expected.TotalTime)
			}
		})
	}
}

func TestGetBestMeal_EdgeCases(t *testing.T) {
	t.Run("Empty dish list", func(t *testing.T) {
		c := chef.NewChef([]chef.Dish{})
		result := c.GetBestMeal()
		
		// Should return zero values
		if result.Entree.Name != "" {
			t.Errorf("Expected empty entree name, got %s", result.Entree.Name)
		}
		if result.TotalTime != 0 {
			t.Errorf("Expected total time 0, got %d", result.TotalTime)
		}
	})
	
	t.Run("Missing dish types", func(t *testing.T) {
		dishes := []chef.Dish{
			{Name: "Only Entree", Type: "entree", PrepTime: 10},
		}
		
		c := chef.NewChef(dishes)
		result := c.GetBestMeal()
		
		// Should handle missing types gracefully
		if result.Entree.Name != "Only Entree" {
			t.Errorf("Expected entree name 'Only Entree', got %s", result.Entree.Name)
		}
		if result.Main.Name != "" {
			t.Errorf("Expected empty main name, got %s", result.Main.Name)
		}
		if result.Dessert.Name != "" {
			t.Errorf("Expected empty dessert name, got %s", result.Dessert.Name)
		}
	})
}

func TestGetBestMeal_TotalTimeCalculation(t *testing.T) {
	dishes := []chef.Dish{
		{Name: "Fast Entree", Type: "entree", PrepTime: 5},
		{Name: "Fast Main", Type: "main", PrepTime: 10},
		{Name: "Fast Dessert", Type: "dessert", PrepTime: 3},
	}
	
	c := chef.NewChef(dishes)
	result := c.GetBestMeal()
	
	expectedTotal := 5 + 10 + 3
	if result.TotalTime != expectedTotal {
		t.Errorf("Total time calculation: got %d, want %d", result.TotalTime, expectedTotal)
	}
	
	// Verify individual times sum to total
	calculatedTotal := result.Entree.PrepTime + result.Main.PrepTime + result.Dessert.PrepTime
	if calculatedTotal != result.TotalTime {
		t.Errorf("Sum of individual times (%d) doesn't match total time (%d)", calculatedTotal, result.TotalTime)
	}
} 
