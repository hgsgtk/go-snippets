# Solution

## Overview

- Data: Stock prices for multiple companies over 7 days
    - e.g., Shopify: [100, 102, 98, 105, 110, 108, 107], Google: [200, 198, 202, 210, 205, 207, 208]
- Goal: Identify the top 3 companies with the highest average stock price for the week.
- Input: Company Stock list: company name and prices
- Output: Top 3 company names

## Requirements

* Is the number of companies always 7? Do you have a plan to scale this solution to more companies?
    * A: Any numbers (thousands or more).
* Are there at least 3 companies in the input? How should we handle the case where there are less than 3 companies?
    * A: Assume that there will always be at least 3 companies in the input.
* Is the stock price's data type always integer?
    * A: The stock prices are of type float64
* What if there are multiple companies with the same average stock price?
    * A:  Break ties by name (lexicographically)
* What if there are multiple companies with the same company name?
    * A: Assume company names are unique.
* Each record must has exactly 7 stock prices.
    * Q: What should we do if there is a new company with less than 7 stock prices?

## Plan

1. Start writing a test case to verify the given example.

```
stocks := []CompanyStock{
    {"Shopify",  {100, 102, 98, 105, 110, 108, 107}},
    {"Google",   {200, 198, 202, 210, 205, 207, 208}},
    {"Amazon",   {150, 148, 152, 153, 151, 149, 150}},
    {"Meta",     {120, 125, 130, 127, 129, 131, 128}},
    {"Netflix",  {300, 305, 310, 315, 320, 325, 330}},
}
```

1. Create the exported function to receive the slice of the CompanyStock struct.
1. Calculate the average stock price for each company over the 7 days.
1. Sort the companies by average stock price in descending order.
1. Return the names of the top 3 companies.

## Review and Optimization

* Average calculation performance:
    * Simple loop: O(n) for time complexity, O(1) for space complexity
    * Alternatives:
        * sum(prices) practically
* Sort performance
    * sort.Slice: quick sort: O(nlogn) for time complexity, O(n) for space complexity
    * Alternatives: 
        * sort.SliceStable: stable merge sort: O(nlogn) for time complexity, O(n) for space complexity
