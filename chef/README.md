# Coding Test Problem: Shortest Path to Complete a Meal

## Problem Description

You are given a list of dishes. Each dish has the following properties:

* name: The name of the dish (string)
* type: The type of the dish, which can be one of "entree", "main", or "dessert"
* prep_time: The preparation time in minutes (integer)

A customer wants to order a complete meal consisting of exactly one entree, one main course, and one dessert. Your task is to select one dish from each category so that the total preparation time is minimized.

## Input

* A list of dish objects, each with:
    * name: string
    * type: string ("entree", "main", or "dessert")
    * prep_time: integer

## Output

* The names of the selected entree, main course, and dessert (in that order)
* The total preparation time (integer)

## Example

Input:

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

Output:

```
["Caesar Salad", "Pasta", "Gelato"], 35
```

## Constraints

* There is at least one dish of each type.
* If there are multiple combinations with the same minimal total time, return any one of them.

## Follow-up

* [ ] How would you modify your solution if some dishes are incompatible with others (for example, certain entrees cannot be paired with certain mains or desserts)?
* [ ] Alternative ways to find the shortest path (performance)