package chef

type Dish struct {
	Name string
	Type string
	PrepTime int
}

type Chef struct {
	Dish []Dish
}

func NewChef(dishes []Dish) *Chef {
	return &Chef{
		Dish: dishes,
	}
}

func (c *Chef) GetBestMeal() ChefRecipe {
	// Track the dish with minimum prep time for each type
	minDishes := make(map[string]Dish)
	
	// Initialize with first dish of each type we encounter
	for _, dish := range c.Dish {
		if _, exists := minDishes[dish.Type]; !exists {
			minDishes[dish.Type] = dish
		} else if dish.PrepTime < minDishes[dish.Type].PrepTime {
			// Update if we find a dish with lower prep time
			minDishes[dish.Type] = dish
		}
	}
	
	// Extract the dishes for each type
	entree := minDishes["entree"]
	main := minDishes["main"]
	dessert := minDishes["dessert"]
	
	// Calculate total prep time
	totalTime := entree.PrepTime + main.PrepTime + dessert.PrepTime
	
	return ChefRecipe{
		Entree:    entree,
		Main:      main,
		Dessert:   dessert,
		TotalTime: totalTime,
	}
}

type ChefRecipe struct {
	Entree Dish
	Main Dish
	Dessert Dish
	TotalTime int
}