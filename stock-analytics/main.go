package main

import (
	"fmt"
	"log"
)

// CompanyStock represents a company and its daily stock prices for a week
type CompanyStock struct {
	Name   string
	Prices []float64 // always 7 daily prices, one for each day of the week
}

func main() {
	// Example usage from the problem statement
	stocks := []CompanyStock{
		{"Shopify", []float64{100, 102, 98, 105, 110, 108, 107}},
		{"Google", []float64{200, 198, 202, 210, 205, 207, 208}},
		{"Amazon", []float64{150, 148, 152, 153, 151, 149, 150}},
		{"Meta", []float64{120, 125, 130, 127, 129, 131, 128}},
		{"Netflix", []float64{300, 305, 310, 315, 320, 325, 330}},
	}

	topCompanies, err := GetTopCompaniesByAveragePrice(stocks)
	if err != nil {
		log.Fatalf("Error getting top companies: %v", err)
	}

	fmt.Println("Top 3 companies by average stock price:")
	for i, company := range topCompanies {
		fmt.Printf("%d. %s\n", i+1, company)
	}
} 
