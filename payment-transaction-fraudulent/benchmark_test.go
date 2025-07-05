package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// generateTestData creates test data with specified parameters
func generateTestData(numTransactions, numPeople, numLocations int) string {
	names := make([]string, numPeople)
	locations := make([]string, numLocations)
	
	// Generate unique names
	for i := 0; i < numPeople; i++ {
		names[i] = fmt.Sprintf("Person%d", i)
	}
	
	// Generate unique locations
	for i := 0; i < numLocations; i++ {
		locations[i] = fmt.Sprintf("Location%d", i)
	}
	
	var transactions []string
	timeCounter := 1
	
	for i := 0; i < numTransactions; i++ {
		name := names[rand.Intn(numPeople)]
		amount := rand.Intn(2000) + 1 // Amount between 1 and 2000
		location := locations[rand.Intn(numLocations)]
		
		transaction := fmt.Sprintf("%s, %d, %d, %s", name, amount, timeCounter, location)
		transactions = append(transactions, transaction)
		
		// Increment time by 1-10 minutes
		timeCounter += rand.Intn(10) + 1
	}
	
	return strings.Join(transactions, "\n")
}

// generateWorstCaseData creates data that triggers the most expensive operations
func generateWorstCaseData(numTransactions int) string {
	var transactions []string
	timeCounter := 1
	
	// Create many transactions for the same person to trigger location change checks
	for i := 0; i < numTransactions; i++ {
		name := "SamePerson" // All transactions for the same person
		amount := 100 // Low amount to avoid high amount rule
		location := fmt.Sprintf("Location%d", i%10) // Cycle through 10 locations
		
		transaction := fmt.Sprintf("%s, %d, %d, %s", name, amount, timeCounter, location)
		transactions = append(transactions, transaction)
		
		// Small time increments to keep transactions within 60-minute window
		timeCounter += rand.Intn(30) + 1
	}
	
	return strings.Join(transactions, "\n")
}

// generateBestCaseData creates data that minimizes expensive operations
func generateBestCaseData(numTransactions int) string {
	var transactions []string
	timeCounter := 1
	
	for i := 0; i < numTransactions; i++ {
		name := fmt.Sprintf("Person%d", i) // Each person has only one transaction
		amount := 100 // Low amount
		location := "SameLocation" // Same location for all
		
		transaction := fmt.Sprintf("%s, %d, %d, %s", name, amount, timeCounter, location)
		transactions = append(transactions, transaction)
		
		timeCounter += 100 // Large time gaps
	}
	
	return strings.Join(transactions, "\n")
}

func BenchmarkFraudDetection_1000Transactions(b *testing.B) {
	// Generate 1000 transactions with 100 people and 20 locations
	data := generateTestData(1000, 100, 20)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := DetectFraudulentTransactions(data)
		if err != nil {
			b.Fatalf("Error in fraud detection: %v", err)
		}
	}
}

func BenchmarkFraudDetection_WorstCase_1000Transactions(b *testing.B) {
	// Generate worst case scenario: all transactions for same person
	data := generateWorstCaseData(1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := DetectFraudulentTransactions(data)
		if err != nil {
			b.Fatalf("Error in fraud detection: %v", err)
		}
	}
}

func BenchmarkFraudDetection_BestCase_1000Transactions(b *testing.B) {
	// Generate best case scenario: each person has only one transaction
	data := generateBestCaseData(1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := DetectFraudulentTransactions(data)
		if err != nil {
			b.Fatalf("Error in fraud detection: %v", err)
		}
	}
}

func BenchmarkFraudDetection_VaryingSizes(b *testing.B) {
	sizes := []int{100, 500, 1000, 2000, 5000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			data := generateTestData(size, size/10, size/50)
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := DetectFraudulentTransactions(data)
				if err != nil {
					b.Fatalf("Error in fraud detection: %v", err)
				}
			}
		})
	}
}

func BenchmarkParseTransactions(b *testing.B) {
	data := generateTestData(1000, 100, 20)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ParseTransactions(data)
		if err != nil {
			b.Fatalf("Error parsing transactions: %v", err)
		}
	}
}

func BenchmarkIsFraudulent_SingleTransaction(b *testing.B) {
	// Create a dataset with one person having multiple transactions
	transactions := []Transaction{
		{Name: "TestPerson", Amount: 100, Time: 10, Location: "Location1"},
		{Name: "TestPerson", Amount: 100, Time: 20, Location: "Location2"},
		{Name: "TestPerson", Amount: 100, Time: 30, Location: "Location1"},
		{Name: "TestPerson", Amount: 100, Time: 40, Location: "Location2"},
		{Name: "TestPerson", Amount: 100, Time: 50, Location: "Location1"},
	}
	
	// Add many other transactions to make the search expensive
	for i := 0; i < 995; i++ {
		transactions = append(transactions, Transaction{
			Name:     fmt.Sprintf("OtherPerson%d", i),
			Amount:   100,
			Time:     100 + i,
			Location: "OtherLocation",
		})
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsFraudulent(transactions[0], transactions)
	}
}

// Performance analysis helper functions
func TestPerformanceAnalysis(t *testing.T) {
	sizes := []int{100, 500, 1000, 2000}
	
	fmt.Println("Performance Analysis:")
	fmt.Println("====================")
	
	for _, size := range sizes {
		data := generateTestData(size, size/10, size/50)
		
		start := time.Now()
		results, err := DetectFraudulentTransactions(data)
		duration := time.Since(start)
		
		if err != nil {
			t.Fatalf("Error in fraud detection: %v", err)
		}
		
		fraudulentCount := 0
		for _, isFraudulent := range results {
			if isFraudulent {
				fraudulentCount++
			}
		}
		
		fmt.Printf("Size: %d transactions\n", size)
		fmt.Printf("  Time: %v\n", duration)
		fmt.Printf("  Fraudulent: %d/%d (%.1f%%)\n", fraudulentCount, size, float64(fraudulentCount)/float64(size)*100)
		fmt.Printf("  Transactions per second: %.0f\n", float64(size)/duration.Seconds())
		fmt.Println()
	}
}

func TestWorstCasePerformance(t *testing.T) {
	fmt.Println("Worst Case Performance Analysis:")
	fmt.Println("===============================")
	
	sizes := []int{100, 500, 1000, 2000}
	
	for _, size := range sizes {
		data := generateWorstCaseData(size)
		
		start := time.Now()
		results, err := DetectFraudulentTransactions(data)
		duration := time.Since(start)
		
		if err != nil {
			t.Fatalf("Error in fraud detection: %v", err)
		}
		
		fraudulentCount := 0
		for _, isFraudulent := range results {
			if isFraudulent {
				fraudulentCount++
			}
		}
		
		fmt.Printf("Size: %d transactions (worst case)\n", size)
		fmt.Printf("  Time: %v\n", duration)
		fmt.Printf("  Fraudulent: %d/%d (%.1f%%)\n", fraudulentCount, size, float64(fraudulentCount)/float64(size)*100)
		fmt.Printf("  Transactions per second: %.0f\n", float64(size)/duration.Seconds())
		fmt.Println()
	}
} 
