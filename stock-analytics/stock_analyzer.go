package main

import (
	"fmt"
	"sort"
)

// CompanyAverage represents a company with its calculated average stock price
type CompanyAverage struct {
	Name    string
	Average float64
}

// GetTopCompaniesByAveragePrice returns the top 3 company names sorted by descending average stock price.
// In case of ties, companies are sorted lexicographically by name.
func GetTopCompaniesByAveragePrice(stocks []CompanyStock) ([]string, error) {
	// Validate input
	if len(stocks) < 3 {
		return nil, fmt.Errorf("at least 3 companies required, got %d", len(stocks))
	}

	// Calculate averages for all companies
	averages := make([]CompanyAverage, 0, len(stocks))
	for _, stock := range stocks {
		if len(stock.Prices) != 7 {
			return nil, fmt.Errorf("company %s must have exactly 7 price points, got %d", stock.Name, len(stock.Prices))
		}

		average := calculateAverage(stock.Prices)
		averages = append(averages, CompanyAverage{
			Name:    stock.Name,
			Average: average,
		})
	}

	// Sort companies by average price (descending), then by name (ascending) for ties
	// Using SliceStable for stable sorting - equal elements maintain their relative order
	sort.SliceStable(averages, func(i, j int) bool {
		if averages[i].Average == averages[j].Average {
			return averages[i].Name < averages[j].Name // lexicographical order for ties
		}
		return averages[i].Average > averages[j].Average // descending order
	})

	// Extract top 3 company names
	result := make([]string, 3)
	for i := 0; i < 3; i++ {
		result[i] = averages[i].Name
	}

	return result, nil
}

// calculateAverage calculates the average of the given price slice
func calculateAverage(prices []float64) float64 {
	if len(prices) == 0 {
		return 0
	}

	sum := 0.0
	for _, price := range prices {
		sum += price
	}
	return sum / float64(len(prices))
} 
