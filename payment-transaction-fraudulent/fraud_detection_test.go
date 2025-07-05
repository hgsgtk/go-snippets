package main

import (
	"testing"
)

func TestFraudDetectionExample(t *testing.T) {
	// Test data from README example
	transactionStrings := []string{
		"Anne, 100, 1, Boston",
		"Anne, 2000, 10, Boston",
		"Bob, 50, 20, Boston",
		"Cindy, 100, 50, New York",
		"Bob, 50, 70, New York",
	}

	// Expected results based on README analysis
	expectedResults := []bool{
		false, // Anne, 100, 1, Boston - Amount ≤ $1,000, same location as next transaction
		true,  // Anne, 2000, 10, Boston - Amount > $1,000
		true, // Bob, 50, 20, Boston - Amount ≤ $1,000, next transaction >= 60 minutes later
		false, // Cindy, 100, 50, New York - Amount ≤ $1,000, no other Cindy transactions
		true,  // Bob, 50, 70, New York - Amount ≤ $1,000, but previous transaction 50 minutes ago at different location
	}

	// Parse all transactions first
	var allTransactions []Transaction
	for _, transactionStr := range transactionStrings {
		transaction, err := ParseTransaction(transactionStr)
		if err != nil {
			t.Fatalf("Failed to parse transaction '%s': %v", transactionStr, err)
		}
		allTransactions = append(allTransactions, transaction)
	}

	// Test each transaction
	for i, transactionStr := range transactionStrings {
		transaction, err := ParseTransaction(transactionStr)
		if err != nil {
			t.Fatalf("Failed to parse transaction '%s': %v", transactionStr, err)
		}
		result := IsFraudulent(transaction, allTransactions)
		expected := expectedResults[i]

		if result != expected {
			t.Errorf("Transaction %d (%s): expected %v, got %v", 
				i+1, transactionStr, expected, result)
		}
	}
}

func TestHighAmountRule(t *testing.T) {
	// Test cases for high amount rule (> $1,000)
	testCases := []struct {
		amount   int
		expected bool
	}{
		{999, false},   // Just under threshold
		{1000, false},  // At threshold
		{1001, true},   // Just over threshold
		{5000, true},   // Well over threshold
	}

	// Create a simple transaction with varying amounts
	baseTransaction := Transaction{
		Name:     "TestUser",
		Time:     100,
		Location: "TestLocation",
	}

	allTransactions := []Transaction{baseTransaction}

	for i, tc := range testCases {
		transaction := baseTransaction
		transaction.Amount = tc.amount
		
		result := IsFraudulent(transaction, allTransactions)
		if result != tc.expected {
			t.Errorf("High amount test %d: amount $%d, expected %v, got %v", 
				i+1, tc.amount, tc.expected, result)
		}
	}
}

func TestLocationChangeRule(t *testing.T) {
	// Test cases for location change rule
	testCases := []struct {
		name     string
		transactions []Transaction
		expected bool
	}{
		{
			name: "Same location within 60 minutes - not fraudulent",
			transactions: []Transaction{
				{Name: "Alice", Amount: 100, Time: 10, Location: "Boston"},
				{Name: "Alice", Amount: 100, Time: 50, Location: "Boston"}, // 40 minutes later, same location
			},
			expected: false,
		},
		{
			name: "Different location within 60 minutes - fraudulent",
			transactions: []Transaction{
				{Name: "Bob", Amount: 100, Time: 10, Location: "Boston"},
				{Name: "Bob", Amount: 100, Time: 50, Location: "New York"}, // 40 minutes later, different location
			},
			expected: true,
		},
		{
			name: "Different location after 60 minutes - not fraudulent",
			transactions: []Transaction{
				{Name: "Charlie", Amount: 100, Time: 10, Location: "Boston"},
				{Name: "Charlie", Amount: 100, Time: 80, Location: "New York"}, // 70 minutes later, different location
			},
			expected: false,
		},
		{
			name: "Single transaction - not fraudulent",
			transactions: []Transaction{
				{Name: "David", Amount: 100, Time: 10, Location: "Boston"},
			},
			expected: false,
		},
	}

	for i, tc := range testCases {
		// Test the last transaction in each test case
		lastTransaction := tc.transactions[len(tc.transactions)-1]
		result := IsFraudulent(lastTransaction, tc.transactions)
		
		if result != tc.expected {
			t.Errorf("Location change test %d (%s): expected %v, got %v", 
				i+1, tc.name, tc.expected, result)
		}
	}
}

func TestEdgeCases(t *testing.T) {
	// Test edge cases
	testCases := []struct {
		name     string
		transactions []Transaction
		expected bool
	}{
		{
			name: "First transaction of a person",
			transactions: []Transaction{
				{Name: "Eve", Amount: 100, Time: 10, Location: "Boston"},
			},
			expected: false,
		},
		{
			name: "Last transaction of a person",
			transactions: []Transaction{
				{Name: "Frank", Amount: 100, Time: 10, Location: "Boston"},
				{Name: "Frank", Amount: 100, Time: 50, Location: "Boston"},
			},
			expected: false,
		},
		{
			name: "Exact 60 minutes difference",
			transactions: []Transaction{
				{Name: "Grace", Amount: 100, Time: 10, Location: "Boston"},
				{Name: "Grace", Amount: 100, Time: 70, Location: "New York"}, // Exactly 60 minutes later
			},
			expected: true, 
		},
	}

	for i, tc := range testCases {
		// Test the last transaction in each test case
		lastTransaction := tc.transactions[len(tc.transactions)-1]
		result := IsFraudulent(lastTransaction, tc.transactions)
		
		if result != tc.expected {
			t.Errorf("Edge case test %d (%s): expected %v, got %v", 
				i+1, tc.name, tc.expected, result)
		}
	}
} 
