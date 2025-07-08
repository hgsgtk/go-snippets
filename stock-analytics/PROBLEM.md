# 🧩 The Problem: Top 3 Companies by Average Stock Price
## Context:
You’re working on an internal analytics tool that processes weekly stock data. You’ve been given stock prices for multiple companies over 7 days, and your goal is to identify the top 3 companies with the highest average stock price for the week.

## Input:
You're given a slice of CompanyStock, each entry representing a company and its daily stock prices for the week:

```go
type CompanyStock struct {
    Name   string
    Prices []float64  // always 7 daily prices, one for each day of the week
}
```

So, your function will receive a slice like this:

```go
stocks := []CompanyStock{
    {"Shopify",  {100, 102, 98, 105, 110, 108, 107}},
    {"Google",   {200, 198, 202, 210, 205, 207, 208}},
    {"Amazon",   {150, 148, 152, 153, 151, 149, 150}},
    {"Meta",     {120, 125, 130, 127, 129, 131, 128}},
    {"Netflix",  {300, 305, 310, 315, 320, 325, 330}},
}
```

## Output:

Return a slice of the top 3 company names, sorted by descending average stock price. In the above example, the expected result would be:

```go
[]string{"Netflix", "Google", "Meta"}
```

## Requirements:

- Calculate the average stock price for each company over the 7 days.
- Sort the companies by average stock price in descending order.
- Return the names of the top 3 companies.
- Assume that there will always be at least 3 companies in the input.
- Feel free to add error handling or sanity checks if you want to simulate real-world robustness.
