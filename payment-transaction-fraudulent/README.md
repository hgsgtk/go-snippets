# Payment Transaction Fraudulent Detection

## Problem Statement

Given a list of transactions consisting of a name, amount, time, and location, determine which transactions are potentially fraudulent.

## Fraud Detection Rules

A transaction is potentially fraudulent if **either** of the following conditions is true:

1. **High Amount Rule**: The transaction amount is greater than `$1,000`
2. **Location Change Rule**: The previous or next transaction *with the same name* occurs at a different location within an hour (`≤ 60 minutes`)

## Input Format

- **Input**: A string containing multiple transactions
- **Format**: Each transaction on a separate line with the format: `<name>, <amount>, <time>, <location>`
- **Time**: Measured in minutes
- **Order**: Transactions are sorted by increasing order by time

### Example Input
```
Anne, 100, 1, Boston
Anne, 2000, 10, Boston
Bob, 50, 20, Boston
Cindy, 100, 50, New York
Bob, 50, 70, New York
```

## Output Format

- **Return**: Boolean indicating whether the transaction is potentially fraudulent
- **True**: Transaction is potentially fraudulent
- **False**: Transaction is not fraudulent

## Example Analysis

Let's analyze the example transactions:

1. **Anne, 100, 1, Boston** → `False`
   - Amount: $100 (≤ $1,000) ✓
   - Next Anne transaction: 9 minutes later, same location (Boston) ✓

2. **Anne, 2000, 10, Boston** → `True`
   - Amount: $2,000 (> $1,000) ❌ (High Amount Rule)

3. **Bob, 50, 20, Boston** → `False`
   - Amount: $50 (≤ $1,000) ✓
   - Next Bob transaction: 50 minutes later, different location (New York) but >= 60 minutes ✓

4. **Cindy, 100, 50, New York** → `False`
   - Amount: $100 (≤ $1,000) ✓
   - No other Cindy transactions ✓

5. **Bob, 50, 70, New York** → `True`
   - Amount: $50 (≤ $1,000) ✓
   - Previous Bob transaction: 50 minutes earlier, different location (Boston) but > 60 minutes ✓
   - Wait, let me recalculate: 70 - 20 = 50 minutes, which is ≤ 60 minutes ❌ (Location Change Rule)

## Implementation Notes

- Transactions are processed in chronological order
- For each transaction, check both fraud detection rules
- Location changes are only considered fraudulent if they occur within 60 minutes
- The same person can have multiple legitimate transactions at different locations if they're more than 60 minutes apart

## Expected Output for Example

```
Anne, 100, 1, Boston → false
Anne, 2000, 10, Boston → true
Bob, 50, 20, Boston → false
Cindy, 100, 50, New York → false
Bob, 50, 70, New York → true
```

## Technical Requirements

- Implement in Go
- Handle edge cases (first/last transactions for a person)
- Efficient processing for large transaction lists
- Clear separation of concerns for fraud detection rules 
