package main

import (
	"fmt"
	"testing"
	"time"
)

// Note: generateTestData, generateWorstCaseData, and generateBestCaseData functions
// are already defined in benchmark_test.go

func BenchmarkOriginalVsOptimized_1000Transactions(b *testing.B) {
	// Generate 1000 transactions with 100 people and 20 locations
	data := generateTestData(1000, 100, 20)
	
	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := DetectFraudulentTransactions(data)
			if err != nil {
				b.Fatalf("Error in original fraud detection: %v", err)
			}
		}
	})
	
	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := DetectFraudulentTransactionsOptimized(data)
			if err != nil {
				b.Fatalf("Error in optimized fraud detection: %v", err)
			}
		}
	})
}

func BenchmarkOriginalVsOptimized_WorstCase_1000Transactions(b *testing.B) {
	// Generate worst case scenario: all transactions for same person
	data := generateWorstCaseData(1000)
	
	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := DetectFraudulentTransactions(data)
			if err != nil {
				b.Fatalf("Error in original fraud detection: %v", err)
			}
		}
	})
	
	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := DetectFraudulentTransactionsOptimized(data)
			if err != nil {
				b.Fatalf("Error in optimized fraud detection: %v", err)
			}
		}
	})
}

func BenchmarkOriginalVsOptimized_BestCase_1000Transactions(b *testing.B) {
	// Generate best case scenario: each person has only one transaction
	data := generateBestCaseData(1000)
	
	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := DetectFraudulentTransactions(data)
			if err != nil {
				b.Fatalf("Error in original fraud detection: %v", err)
			}
		}
	})
	
	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := DetectFraudulentTransactionsOptimized(data)
			if err != nil {
				b.Fatalf("Error in optimized fraud detection: %v", err)
			}
		}
	})
}

func BenchmarkOriginalVsOptimized_VaryingSizes(b *testing.B) {
	sizes := []int{100, 500, 1000, 2000, 5000}
	
	for _, size := range sizes {
		data := generateTestData(size, size/10, size/50)
		
		b.Run(fmt.Sprintf("Original_Size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := DetectFraudulentTransactions(data)
				if err != nil {
					b.Fatalf("Error in original fraud detection: %v", err)
				}
			}
		})
		
		b.Run(fmt.Sprintf("Optimized_Size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := DetectFraudulentTransactionsOptimized(data)
				if err != nil {
					b.Fatalf("Error in optimized fraud detection: %v", err)
				}
			}
		})
	}
}

// Performance comparison test
func TestPerformanceComparison(t *testing.T) {
	sizes := []int{100, 500, 1000, 2000}
	
	fmt.Println("Performance Comparison: Original vs Optimized")
	fmt.Println("=============================================")
	
	for _, size := range sizes {
		data := generateTestData(size, size/10, size/50)
		
		// Test original implementation
		start := time.Now()
		originalResults, err := DetectFraudulentTransactions(data)
		originalTime := time.Since(start)
		if err != nil {
			t.Fatalf("Original implementation failed: %v", err)
		}
		
		// Test optimized implementation
		start = time.Now()
		optimizedResults, err := DetectFraudulentTransactionsOptimized(data)
		optimizedTime := time.Since(start)
		if err != nil {
			t.Fatalf("Optimized implementation failed: %v", err)
		}
		
		// Verify results match
		resultsMatch := true
		for i, original := range originalResults {
			if original != optimizedResults[i] {
				resultsMatch = false
				t.Errorf("Result mismatch at index %d: original=%v, optimized=%v", i, original, optimizedResults[i])
			}
		}
		
		// Calculate improvement
		improvement := float64(originalTime) / float64(optimizedTime)
		
		fmt.Printf("Size: %d transactions\n", size)
		fmt.Printf("  Original time: %v\n", originalTime)
		fmt.Printf("  Optimized time: %v\n", optimizedTime)
		fmt.Printf("  Performance improvement: %.2fx\n", improvement)
		fmt.Printf("  Results match: %v\n", resultsMatch)
		fmt.Println()
	}
}

func TestWorstCasePerformanceComparison(t *testing.T) {
	sizes := []int{100, 500, 1000, 2000}
	
	fmt.Println("Worst Case Performance Comparison: Original vs Optimized")
	fmt.Println("=======================================================")
	
	for _, size := range sizes {
		data := generateWorstCaseData(size)
		
		// Test original implementation
		start := time.Now()
		originalResults, err := DetectFraudulentTransactions(data)
		originalTime := time.Since(start)
		if err != nil {
			t.Fatalf("Original implementation failed: %v", err)
		}
		
		// Test optimized implementation
		start = time.Now()
		optimizedResults, err := DetectFraudulentTransactionsOptimized(data)
		optimizedTime := time.Since(start)
		if err != nil {
			t.Fatalf("Optimized implementation failed: %v", err)
		}
		
		// Verify results match
		resultsMatch := true
		for i, original := range originalResults {
			if original != optimizedResults[i] {
				resultsMatch = false
				t.Errorf("Result mismatch at index %d: original=%v, optimized=%v", i, original, optimizedResults[i])
			}
		}
		
		// Calculate improvement
		improvement := float64(originalTime) / float64(optimizedTime)
		
		fmt.Printf("Size: %d transactions (worst case)\n", size)
		fmt.Printf("  Original time: %v\n", originalTime)
		fmt.Printf("  Optimized time: %v\n", optimizedTime)
		fmt.Printf("  Performance improvement: %.2fx\n", improvement)
		fmt.Printf("  Results match: %v\n", resultsMatch)
		fmt.Println()
	}
} 
