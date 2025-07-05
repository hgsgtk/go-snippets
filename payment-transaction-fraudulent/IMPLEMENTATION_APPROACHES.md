# Implementation Approaches for Fraud Detection

## Problem Analysis

The fraud detection system needs to:
1. Parse transaction strings into structured data
2. Check if amount > $1,000 (High Amount Rule)
3. Check if previous/next transaction with same name occurs at different location within 60 minutes (Location Change Rule)

## Approach 1: Simple Linear Search (Brute Force)

### Implementation Strategy
- Parse all transactions into a slice of Transaction structs
- For each transaction, scan through all other transactions to find same-name transactions
- Check time differences and location changes

### Code Structure
```go
func IsFraudulent(transaction Transaction, allTransactions []Transaction) bool {
    // Check high amount rule
    if transaction.Amount > 1000 {
        return true
    }
    
    // Check location change rule
    for _, other := range allTransactions {
        if other.Name == transaction.Name && other != transaction {
            timeDiff := abs(transaction.Time - other.Time)
            if timeDiff <= 60 && other.Location != transaction.Location {
                return true
            }
        }
    }
    return false
}
```

### Pros
- Simple to implement and understand
- No complex data structures
- Easy to debug and test

### Cons
- O(n²) time complexity for each transaction
- Poor performance with large datasets
- Redundant calculations

### Best For
- Small datasets (< 1000 transactions)
- Prototyping and initial development
- Educational purposes

---

## Approach 2: Pre-sorted by Name and Time

### Implementation Strategy
- Sort transactions by name first, then by time
- Group transactions by name
- For each transaction, only check adjacent transactions in the same group

### Code Structure
```go
func IsFraudulent(transaction Transaction, allTransactions []Transaction) bool {
    // Check high amount rule
    if transaction.Amount > 1000 {
        return true
    }
    
    // Find transactions with same name
    sameNameTransactions := filterByName(allTransactions, transaction.Name)
    
    // Check adjacent transactions within 60 minutes
    for _, other := range sameNameTransactions {
        if other != transaction {
            timeDiff := abs(transaction.Time - other.Time)
            if timeDiff <= 60 && other.Location != transaction.Location {
                return true
            }
        }
    }
    return false
}
```

### Pros
- Better performance than brute force
- Logical grouping by person
- Easier to understand fraud patterns

### Cons
- Still O(n) per transaction in worst case
- Requires sorting/grouping step
- Memory overhead for grouping

### Best For
- Medium datasets (1000-10000 transactions)
- When you need to analyze patterns by person

---

## Approach 3: Hash Map with Time Windows

### Implementation Strategy
- Use a map to group transactions by name
- For each person, maintain a sliding window of recent transactions
- Check only transactions within the 60-minute window

### Code Structure
```go
type PersonTransactions struct {
    recent []Transaction // transactions within last 60 minutes
}

func IsFraudulent(transaction Transaction, transactionMap map[string]*PersonTransactions) bool {
    // Check high amount rule
    if transaction.Amount > 1000 {
        return true
    }
    
    person := transactionMap[transaction.Name]
    if person == nil {
        return false
    }
    
    // Check recent transactions for location changes
    for _, recent := range person.recent {
        if recent.Location != transaction.Location {
            return true
        }
    }
    return false
}
```

### Pros
- O(1) average case for location checks
- Efficient memory usage
- Good for real-time processing

### Cons
- More complex implementation
- Requires maintaining time windows
- Potential for bugs in window management

### Best For
- Real-time transaction processing
- Large datasets with many transactions per person

---

## Approach 4: Binary Search with Sorted Arrays

### Implementation Strategy
- Maintain sorted arrays of transactions per person
- Use binary search to find transactions within time windows
- Efficient range queries for location changes

### Code Structure
```go
type PersonHistory struct {
    transactions []Transaction // sorted by time
}

func IsFraudulent(transaction Transaction, personMap map[string]*PersonHistory) bool {
    // Check high amount rule
    if transaction.Amount > 1000 {
        return true
    }
    
    person := personMap[transaction.Name]
    if person == nil {
        return false
    }
    
    // Binary search for transactions within 60 minutes
    startTime := transaction.Time - 60
    endTime := transaction.Time + 60
    
    startIdx := binarySearch(person.transactions, startTime)
    endIdx := binarySearch(person.transactions, endTime)
    
    // Check for location changes in range
    for i := startIdx; i <= endIdx; i++ {
        if person.transactions[i].Location != transaction.Location {
            return true
        }
    }
    return false
}
```

### Pros
- O(log n) for finding time windows
- Efficient for large transaction histories
- Good for batch processing

### Cons
- Complex binary search implementation
- Requires maintaining sorted arrays
- Memory overhead for duplicate data

### Best For
- Large datasets with long transaction histories
- Batch processing scenarios

---

## Approach 5: Sliding Window with Circular Buffer

### Implementation Strategy
- Use a circular buffer to maintain recent transactions per person
- Fixed memory usage regardless of transaction volume
- Efficient for real-time processing

### Code Structure
```go
type CircularBuffer struct {
    buffer []Transaction
    head   int
    size   int
    maxSize int
}

func IsFraudulent(transaction Transaction, buffers map[string]*CircularBuffer) bool {
    // Check high amount rule
    if transaction.Amount > 1000 {
        return true
    }
    
    buffer := buffers[transaction.Name]
    if buffer == nil {
        return false
    }
    
    // Check all transactions in buffer for location changes
    for i := 0; i < buffer.size; i++ {
        idx := (buffer.head + i) % buffer.maxSize
        recent := buffer.buffer[idx]
        if abs(transaction.Time - recent.Time) <= 60 && 
           recent.Location != transaction.Location {
            return true
        }
    }
    return false
}
```

### Pros
- Fixed memory usage
- Good for real-time processing
- Predictable performance

### Cons
- Limited history (only recent transactions)
- May miss fraud patterns beyond buffer size
- Complex buffer management

### Best For
- Real-time systems with memory constraints
- When only recent history matters

---

## Approach 6: Database-like Indexing

### Implementation Strategy
- Create indexes on name, time, and location
- Use efficient data structures for range queries
- Similar to database query optimization

### Code Structure
```go
type TransactionIndex struct {
    byName map[string][]int // name -> transaction indices
    byTime []int            // sorted transaction indices by time
}

func IsFraudulent(transaction Transaction, index *TransactionIndex) bool {
    // Check high amount rule
    if transaction.Amount > 1000 {
        return true
    }
    
    // Use index to find relevant transactions efficiently
    nameIndices := index.byName[transaction.Name]
    relevantIndices := findInTimeRange(nameIndices, transaction.Time-60, transaction.Time+60)
    
    for _, idx := range relevantIndices {
        if allTransactions[idx].Location != transaction.Location {
            return true
        }
    }
    return false
}
```

### Pros
- Very efficient for large datasets
- Good for complex queries
- Scalable architecture

### Cons
- Complex implementation
- Memory overhead for indexes
- Overkill for simple use cases

### Best For
- Very large datasets (>100k transactions)
- Complex fraud detection rules
- Production systems

---

## Recommended Approach Selection

### For Learning/Prototyping: Approach 1 (Simple Linear Search)
- Start here to understand the problem
- Easy to implement and debug
- Good foundation for more complex approaches

### For Medium Datasets: Approach 2 (Pre-sorted by Name)
- Good balance of simplicity and performance
- Easy to understand and maintain
- Suitable for most practical use cases

### For Production Systems: Approach 3 (Hash Map with Time Windows)
- Best performance for real-time processing
- Reasonable complexity
- Good for high-volume transaction processing

### For Large-scale Systems: Approach 6 (Database-like Indexing)
- Maximum performance and scalability
- Complex but very efficient
- Best for enterprise-level fraud detection

## Implementation Priority

1. **Start with Approach 1** - Get it working correctly
2. **Profile performance** - Identify bottlenecks
3. **Upgrade to Approach 2 or 3** - Based on performance needs
4. **Consider Approach 6** - Only if dealing with massive datasets

## Testing Strategy

- Use the existing test file to validate all approaches
- Add performance benchmarks for comparison
- Test with various dataset sizes and patterns 
