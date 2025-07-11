# Coding Test Problem: Implement a Grep-like Feature

## Problem Description

You are tasked with implementing a simplified version of the `grep` command-line utility. Your program should search through a given text input and return all lines that contain a specified search pattern.

## Input

1. **Pattern**: A string representing the search pattern.
2. **Text**: Multiple lines of text where the search will be performed.

## Output

- All lines from the input text that contain the search pattern.
- The matching lines should be output in the order they appear in the input.
- If no lines match, output nothing.

## Requirements

- The search should be **case-sensitive**.
- The pattern can appear anywhere within the line.
- The program should handle any number of input lines.
- Do not use any built-in or external library functions that directly perform pattern matching or searching (e.g., no regex libraries or built-in `grep` commands).
- The solution should be language-agnostic and focus on the algorithmic approach.

## Example

### Input

```
Pattern: "apple"
Text:
I have an apple.
This is a banana.
Apple pie is delicious.
An apple a day keeps the doctor away.
```

### Output

```
I have an apple.
An apple a day keeps the doctor away.
```

## Additional Notes

- The problem tests string searching and line filtering skills.
- Candidates should demonstrate their understanding of substring search algorithms (e.g., naive search, Knuth-Morris-Pratt, or other efficient methods).
- The problem encourages handling input/output and string processing without relying on language-specific features.
