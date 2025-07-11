# Implementation Plan

## User Story

```go
list := []Dish{
  { name: "Bruschetta", type: "entree", prep_time: 10 },
  { name: "Caesar Salad", type: "entree", prep_time: 8 },
  { name: "Steak", type: "main", prep_time: 25 },
  { name: "Pasta", type: "main", prep_time: 20 },
  { name: "Tiramisu", type: "dessert", prep_time: 12 },
  { name: "Gelato", type: "dessert", prep_time: 7 }
}

c := NewChef(list)
output := c.GetBestMeal()

// Output: ["Caesar Salad", "Pasta", "Gelato"], 35
```

## Functionality

- [x] Accept a list of dishes
    * Constraints: At least one dish of each type.
    * name: string (unique)
- [x] Return the best meal
    * Constraints: No order requirement when two dishes have the same prep time.
- [x] Return the total preparation time

## Interface

* Library
* Input: list of dishes

```
[
  { name: "Bruschetta", type: "entree", prep_time: 10 },
  { name: "Caesar Salad", type: "entree", prep_time: 8 },
  { name: "Steak", type: "main", prep_time: 25 },
  { name: "Pasta", type: "main", prep_time: 20 },
  { name: "Tiramisu", type: "dessert", prep_time: 12 },
  { name: "Gelato", type: "dessert", prep_time: 7 }
]
```

## Extra bonus implementation

#### some dishes are incompatible with others

1. Add a field to the Dish struct to store the incompatible dishes.

```
type Dish struct {
	Name string
	Type string
	PrepTime int
	IncompatibleDishes []string
}
```

2. Check the incompatibilities.

```go
  // Check if the dish is incompatible with the current min dish
  if !c.IsIncompatible(dish, minDishes[dish.Type]) {
    minDishes[dish.Type] = dish
  }
```
