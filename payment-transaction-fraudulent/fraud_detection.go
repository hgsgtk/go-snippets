package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Transaction represents a payment transaction
type Transaction struct {
	Name     string
	Amount   int
	Time     int
	Location string
}

// ParseTransaction parses a transaction string into a Transaction struct
// Format: "name, amount, time, location"
func ParseTransaction(transactionStr string) (Transaction, error) {
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

// ParseTransactions parses a string containing multiple transactions
// Each transaction should be on a separate line
func ParseTransactions(input string) ([]Transaction, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	transactions := make([]Transaction, 0, len(lines))

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // Skip empty lines
		}

		transaction, err := ParseTransaction(line)
		if err != nil {
			return nil, fmt.Errorf("line %d: %v", i+1, err)
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// IsFraudulent determines if a transaction is potentially fraudulent
// Rules:
// 1. Amount > $1,000
// 2. Previous or next transaction with same name at different location within 60 minutes
func IsFraudulent(transaction Transaction, allTransactions []Transaction) bool {
	// Rule 1: Check if amount is greater than $1,000
	if transaction.Amount > 1000 {
		return true
	}

	// Find all transactions with the same name, sorted by time
	var sameName []Transaction
	for _, other := range allTransactions {
		if other.Name == transaction.Name {
			sameName = append(sameName, other)
		}
	}

	// Find the index of the current transaction in sameName
	idx := -1
	for i, t := range sameName {
		if t == transaction {
			idx = i
			break
		}
	}

	// Check previous transaction with same name
	if idx > 0 {
		prev := sameName[idx-1]
		if abs(transaction.Time-prev.Time) <= 60 && prev.Location != transaction.Location {
			return true
		}
	}

	// Check next transaction with same name
	if idx >= 0 && idx < len(sameName)-1 {
		next := sameName[idx+1]
		if abs(transaction.Time-next.Time) <= 60 && next.Location != transaction.Location {
			return true
		}
	}

	return false
}

// DetectFraudulentTransactions processes all transactions and returns which ones are fraudulent
func DetectFraudulentTransactions(input string) ([]bool, error) {
	transactions, err := ParseTransactions(input)
	if err != nil {
		return nil, err
	}

	results := make([]bool, len(transactions))
	for i, transaction := range transactions {
		results[i] = IsFraudulent(transaction, transactions)
	}

	return results, nil
}

// FormatResults formats the fraud detection results for display
func FormatResults(transactionStrings []string, results []bool) string {
	var output strings.Builder
	output.WriteString("Fraud Detection Results:\n")
	output.WriteString("========================\n")

	for i, transactionStr := range transactionStrings {
		status := "LEGITIMATE"
		if results[i] {
			status = "FRAUDULENT"
		}
		output.WriteString(fmt.Sprintf("%s → %s\n", transactionStr, status))
	}

	return output.String()
} 
