# Stock Analytics - Top 3 Companies by Average Stock Price

A Go implementation for analyzing weekly stock data and identifying the top 3 companies with the highest average stock price.

## Problem Statement

Given stock prices for multiple companies over 7 days, identify the top 3 companies with the highest average stock price for the week.

## Features

- ✅ Calculate average stock prices for multiple companies
- ✅ Sort companies by average price in descending order
- ✅ Handle tie-breaking using lexicographical company name ordering
- ✅ Comprehensive input validation
- ✅ Extensive test coverage
- ✅ Performance benchmarking
- ✅ Error handling for edge cases

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    stocks := []CompanyStock{
        {"Shopify", []float64{100, 102, 98, 105, 110, 108, 107}},
        {"Google", []float64{200, 198, 202, 210, 205, 207, 208}},
        {"Amazon", []float64{150, 148, 152, 153, 151, 149, 150}},
        {"Meta", []float64{120, 125, 130, 127, 129, 131, 128}},
        {"Netflix", []float64{300, 305, 310, 315, 320, 325, 330}},
    }

    topCompanies, err := GetTopCompaniesByAveragePrice(stocks)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    fmt.Println("Top 3 companies:", topCompanies)
    // Output: [Netflix Google Amazon]
}
```

### Running the Example

```bash
go run .
```

### Running Tests

```bash
# Run all tests
go test -v

# Run benchmarks
go test -bench=. -benchmem
```

## API Reference

### Types

```go
type CompanyStock struct {
    Name   string
    Prices []float64 // always 7 daily prices, one for each day of the week
}
```

### Functions

#### `GetTopCompaniesByAveragePrice(stocks []CompanyStock) ([]string, error)`

Returns the top 3 company names sorted by descending average stock price.

**Parameters:**
- `stocks`: Slice of CompanyStock containing company data

**Returns:**
- `[]string`: Top 3 company names in order
- `error`: Error if validation fails

**Error Cases:**
- Less than 3 companies provided
- Company with incorrect number of price points (not 7)
- Empty input

## Requirements

- **Input Validation**: At least 3 companies required
- **Price Points**: Each company must have exactly 7 daily prices
- **Data Type**: Stock prices are float64
- **Tie Breaking**: Companies with same average are sorted by name (lexicographically)
- **Unique Names**: Company names are assumed to be unique

## Performance

Benchmark results for 1000 companies:
```
BenchmarkGetTopCompaniesByAveragePrice-8   85154   13814 ns/op   24720 B/op   5 allocs/op
```

## Test Coverage

The implementation includes comprehensive tests covering:

- ✅ Example from problem statement
- ✅ Exactly 3 companies
- ✅ Tie-breaking scenarios
- ✅ Error cases (insufficient companies, wrong price count)
- ✅ Edge cases (empty input)
- ✅ Average calculation accuracy
- ✅ Performance benchmarking

## Project Structure

```
stock-analytics/
├── main.go                    # Example usage and entry point
├── stock_analyzer.go          # Core business logic
├── stock_analyzer_test.go     # Comprehensive test suite
├── go.mod                     # Go module file
├── README.md                  # This documentation
├── PROBLEM.md                 # Original problem statement
└── SOLUTION.md                # Solution approach and requirements
```

## Algorithm

1. **Validation**: Check for minimum 3 companies and correct price count
2. **Average Calculation**: Compute average stock price for each company
3. **Sorting**: Sort companies by average price (descending), then by name (ascending) for ties
4. **Extraction**: Return top 3 company names

## Future Enhancements

- Support for variable number of days (not just 7)
- Configurable top N companies (not just top 3)
- Additional aggregation methods (median, weighted average)
- Caching for repeated calculations
- Database integration for persistent storage
- REST API interface 
