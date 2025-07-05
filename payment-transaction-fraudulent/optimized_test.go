package main

import (
	"testing"
)

func TestOptimizedImplementation(t *testing.T) {
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
		true,  // Bob, 50, 20, Boston - Amount ≤ $1,000, next transaction >= 60 minutes later
		false, // Cindy, 100, 50, New York - Amount ≤ $1,000, no other Cindy transactions
		true,  // Bob, 50, 70, New York - Amount ≤ $1,000, but previous transaction 50 minutes ago at different location
	}

	// Test optimized implementation
	input := "Anne, 100, 1, Boston\nAnne, 2000, 10, Boston\nBob, 50, 20, Boston\nCindy, 100, 50, New York\nBob, 50, 70, New York"
	
	results, err := DetectFraudulentTransactionsOptimized(input)
	if err != nil {
		t.Fatalf("Optimized implementation failed: %v", err)
	}

	// Verify results
	for i, expected := range expectedResults {
		if results[i] != expected {
			t.Errorf("Transaction %d (%s): expected %v, got %v", 
				i+1, transactionStrings[i], expected, results[i])
		}
	}
}

func TestOptimizedVsOriginal(t *testing.T) {
	// Test that optimized implementation produces same results as original
	testCases := []string{
		// Simple case
		"Alice, 100, 10, Boston\nAlice, 200, 50, New York",
		
		// Multiple people
		"Bob, 100, 10, Boston\nCharlie, 200, 20, New York\nBob, 300, 30, Boston",
		
		// High amount case
		"David, 1500, 10, Boston\nDavid, 100, 20, Boston",
		
		// Edge case - exact 60 minutes
		"Eve, 100, 10, Boston\nEve, 100, 70, New York",
	}

	for i, input := range testCases {
		// Get original results
		originalResults, err := DetectFraudulentTransactions(input)
		if err != nil {
			t.Fatalf("Original implementation failed for test case %d: %v", i+1, err)
		}

		// Get optimized results
		optimizedResults, err := DetectFraudulentTransactionsOptimized(input)
		if err != nil {
			t.Fatalf("Optimized implementation failed for test case %d: %v", i+1, err)
		}

		// Compare results
		if len(originalResults) != len(optimizedResults) {
			t.Errorf("Test case %d: result length mismatch - original: %d, optimized: %d", 
				i+1, len(originalResults), len(optimizedResults))
			continue
		}

		for j, original := range originalResults {
			if original != optimizedResults[j] {
				t.Errorf("Test case %d, transaction %d: result mismatch - original: %v, optimized: %v", 
					i+1, j+1, original, optimizedResults[j])
			}
		}
	}
}

func TestGroupTransactionsByName(t *testing.T) {
	transactions := []Transaction{
		{Name: "Alice", Amount: 100, Time: 20, Location: "Boston"},
		{Name: "Bob", Amount: 200, Time: 10, Location: "New York"},
		{Name: "Alice", Amount: 300, Time: 30, Location: "Boston"},
		{Name: "Charlie", Amount: 400, Time: 15, Location: "Chicago"},
	}

	groups := GroupTransactionsByName(transactions)

	// Check that all transactions are grouped
	if len(groups) != 3 {
		t.Errorf("Expected 3 groups, got %d", len(groups))
	}

	// Check Alice's transactions are sorted by time
	aliceGroup := groups["Alice"]
	if aliceGroup == nil {
		t.Fatal("Alice group not found")
	}
	if len(aliceGroup.Transactions) != 2 {
		t.Errorf("Expected 2 transactions for Alice, got %d", len(aliceGroup.Transactions))
	}
	if aliceGroup.Transactions[0].Time != 20 || aliceGroup.Transactions[1].Time != 30 {
		t.Errorf("Alice's transactions not sorted by time: %v", aliceGroup.Transactions)
	}

	// Check Bob's transactions
	bobGroup := groups["Bob"]
	if bobGroup == nil {
		t.Fatal("Bob group not found")
	}
	if len(bobGroup.Transactions) != 1 {
		t.Errorf("Expected 1 transaction for Bob, got %d", len(bobGroup.Transactions))
	}

	// Check Charlie's transactions
	charlieGroup := groups["Charlie"]
	if charlieGroup == nil {
		t.Fatal("Charlie group not found")
	}
	if len(charlieGroup.Transactions) != 1 {
		t.Errorf("Expected 1 transaction for Charlie, got %d", len(charlieGroup.Transactions))
	}
} 
