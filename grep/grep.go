package grep

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// isRegexPattern checks if the pattern contains regex metacharacters
func isRegexPattern(pattern string) bool {
	regexChars := []string{".", "*", "+", "?", "^", "$", "[", "]", "(", ")", "|", "\\", "{", "}"}
	for _, char := range regexChars {
		if strings.Contains(pattern, char) {
			return true
		}
	}
	return false
}

// Grep searches for lines containing the pattern in the given text.
// It returns all matching lines in the order they appear in the input.
// The pattern is automatically detected as regex if it contains regex metacharacters.
func Grep(pattern string, text io.Reader) ([]string, error) {
	if pattern == "" {
		return nil, fmt.Errorf("pattern cannot be empty")
	}

	var matches []string
	scanner := bufio.NewScanner(text)

	// Auto-detect if pattern is regex
	useRegex := isRegexPattern(pattern)

	if useRegex {
		// Compile regex pattern
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("invalid regex pattern: %w", err)
		}

		// Search using regex
		for scanner.Scan() {
			line := scanner.Text()
			if re.MatchString(line) {
				matches = append(matches, line)
			}
		}
	} else {
		// Search using plain text
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, pattern) {
				matches = append(matches, line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}

	return matches, nil
} 