package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Note: Transaction struct is already defined in fraud_detection.go

// PersonTransactions represents all transactions for a single person
type PersonTransactions struct {
	Name         string
	Transactions []Transaction
}

// ParseTransaction parses a transaction string into a Transaction struct
// Format: "name, amount, time, location"
func ParseTransactionOptimized(transactionStr string) (Transaction, error) {
	parts := strings.Split(transactionStr, ",")
	if len(parts) != 4 {
		return Transaction{}, fmt.Errorf("invalid transaction format: expected 4 parts, got %d", len(parts))
	}

	// Trim whitespace from each part
	name := strings.TrimSpace(parts[0])
	amountStr := strings.TrimSpace(parts[1])
	timeStr := strings.TrimSpace(parts[2])
	location := strings.TrimSpace(parts[3])

	// Parse amount
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return Transaction{}, fmt.Errorf("invalid amount '%s': %v", amountStr, err)
	}

	// Parse time
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return Transaction{}, fmt.Errorf("invalid time '%s': %v", timeStr, err)
	}

	return Transaction{
		Name:     name,
		Amount:   amount,
		Time:     time,
		Location: location,
	}, nil
}

// ParseTransactionsOptimized parses a string containing multiple transactions
// Each transaction should be on a separate line
func ParseTransactionsOptimized(input string) ([]Transaction, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	transactions := make([]Transaction, 0, len(lines))

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // Skip empty lines
		}

		transaction, err := ParseTransactionOptimized(line)
		if err != nil {
			return nil, fmt.Errorf("line %d: %v", i+1, err)
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// Note: abs function is already defined in fraud_detection.go

// GroupTransactionsByName groups transactions by person name and sorts each group by time
func GroupTransactionsByName(transactions []Transaction) map[string]*PersonTransactions {
	groups := make(map[string]*PersonTransactions)

	// Group transactions by name
	for _, transaction := range transactions {
		if groups[transaction.Name] == nil {
			groups[transaction.Name] = &PersonTransactions{
				Name:         transaction.Name,
				Transactions: make([]Transaction, 0, 10), // Pre-allocate with reasonable capacity
			}
		}
		groups[transaction.Name].Transactions = append(groups[transaction.Name].Transactions, transaction)
	}

	// Sort each group by time
	for _, group := range groups {
		sort.Slice(group.Transactions, func(i, j int) bool {
			return group.Transactions[i].Time < group.Transactions[j].Time
		})
	}

	return groups
}

// IsFraudulentOptimized determines if a transaction is potentially fraudulent using the optimized approach
// Rules:
// 1. Amount > $1,000
// 2. Previous or next transaction with same name at different location within 60 minutes
func IsFraudulentOptimized(transaction Transaction, groups map[string]*PersonTransactions) bool {
	// Rule 1: Check if amount is greater than $1,000
	if transaction.Amount > 1000 {
		return true
	}

	// Get the person's transaction group
	personGroup := groups[transaction.Name]
	if personGroup == nil {
		return false // No other transactions for this person
	}

	// Find the index of the current transaction in the sorted group
	idx := -1
	for i, t := range personGroup.Transactions {
		if t == transaction {
			idx = i
			break
		}
	}

	if idx == -1 {
		return false // Transaction not found in group (shouldn't happen)
	}

	// Check previous transaction with same name
	if idx > 0 {
		prev := personGroup.Transactions[idx-1]
		if abs(transaction.Time-prev.Time) <= 60 && prev.Location != transaction.Location {
			return true
		}
	}

	// Check next transaction with same name
	if idx < len(personGroup.Transactions)-1 {
		next := personGroup.Transactions[idx+1]
		if abs(transaction.Time-next.Time) <= 60 && next.Location != transaction.Location {
			return true
		}
	}

	return false
}

// DetectFraudulentTransactionsOptimized processes all transactions using the optimized approach
func DetectFraudulentTransactionsOptimized(input string) ([]bool, error) {
	transactions, err := ParseTransactionsOptimized(input)
	if err != nil {
		return nil, err
	}

	// Group transactions by name and sort each group by time
	groups := GroupTransactionsByName(transactions)

	results := make([]bool, len(transactions))
	for i, transaction := range transactions {
		results[i] = IsFraudulentOptimized(transaction, groups)
	}

	return results, nil
}

// FormatResultsOptimized formats the fraud detection results for display
func FormatResultsOptimized(transactionStrings []string, results []bool) string {
	var output strings.Builder
	output.WriteString("Fraud Detection Results (Optimized):\n")
	output.WriteString("====================================\n")

	for i, transactionStr := range transactionStrings {
		status := "LEGITIMATE"
		if results[i] {
			status = "FRAUDULENT"
		}
		output.WriteString(fmt.Sprintf("%s → %s\n", transactionStr, status))
	}

	return output.String()
}

// Performance comparison helper function
func ComparePerformance(input string) error {
	fmt.Println("Performance Comparison:")
	fmt.Println("=======================")

	// Test original implementation
	start := time.Now()
	originalResults, err := DetectFraudulentTransactions(input)
	originalTime := time.Since(start)
	if err != nil {
		return fmt.Errorf("original implementation failed: %v", err)
	}

	// Test optimized implementation
	start = time.Now()
	optimizedResults, err := DetectFraudulentTransactionsOptimized(input)
	optimizedTime := time.Since(start)
	if err != nil {
		return fmt.Errorf("optimized implementation failed: %v", err)
	}

	// Compare results
	resultsMatch := true
	for i, original := range originalResults {
		if original != optimizedResults[i] {
			resultsMatch = false
			fmt.Printf("Result mismatch at index %d: original=%v, optimized=%v\n", i, original, optimizedResults[i])
		}
	}

	fmt.Printf("Results match: %v\n", resultsMatch)
	fmt.Printf("Original implementation time: %v\n", originalTime)
	fmt.Printf("Optimized implementation time: %v\n", optimizedTime)
	fmt.Printf("Performance improvement: %.2fx\n", float64(originalTime)/float64(optimizedTime))

	return nil
} 
