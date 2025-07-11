package grep

import (
	"strings"
	"testing"
)

func TestGrep_PlainText(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		input    string
		expected []string
	}{
		{
			name:     "simple match",
			pattern:  "apple",
			input:    "I have an apple.\nThis is a banana.\nAn apple a day keeps the doctor away.",
			expected: []string{"I have an apple.", "An apple a day keeps the doctor away."},
		},
		{
			name:     "case sensitive",
			pattern:  "Apple",
			input:    "I have an apple.\nApple pie is delicious.\nAn apple a day keeps the doctor away.",
			expected: []string{"Apple pie is delicious."},
		},
		{
			name:     "no match",
			pattern:  "orange",
			input:    "I have an apple.\nThis is a banana.\nAn apple a day keeps the doctor away.",
			expected: []string{},
		},
		{
			name:     "exact word match",
			pattern:  "apple",
			input:    "apple\napples\npineapple\nI have an apple.",
			expected: []string{"apple", "apples", "pineapple", "I have an apple."},
		},
		{
			name:     "empty input",
			pattern:  "apple",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single line",
			pattern:  "test",
			input:    "This is a test line.",
			expected: []string{"This is a test line."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			result, err := Grep(tt.pattern, reader)
			if err != nil {
				t.Errorf("Grep() error = %v", err)
				return
			}
			if !compareSlices(result, tt.expected) {
				t.Errorf("Grep() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGrep_Regex(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		input    string
		expected []string
	}{
		{
			name:     "start of line",
			pattern:  "^apple",
			input:    "apple pie\nI have an apple.\napple juice\nAn apple a day",
			expected: []string{"apple pie", "apple juice"},
		},
		{
			name:     "end of line",
			pattern:  "apple$",
			input:    "I have an apple\napple pie\nred apple\nAn apple a day",
			expected: []string{"I have an apple", "red apple"},
		},
		{
			name:     "alternation",
			pattern:  "apple|banana",
			input:    "I have an apple.\nThis is a banana.\nApple pie is delicious.\nAn apple a day keeps the doctor away.",
			expected: []string{"I have an apple.", "This is a banana.", "An apple a day keeps the doctor away."},
		},
		{
			name:     "character class",
			pattern:  "[Aa]pple",
			input:    "I have an apple.\nApple pie is delicious.\nAn apple a day keeps the doctor away.",
			expected: []string{"I have an apple.", "Apple pie is delicious.", "An apple a day keeps the doctor away."},
		},
		{
			name:     "quantifier",
			pattern:  "appl*e",
			input:    "apple\napplle\napplle\nbanana",
			expected: []string{"apple", "applle", "applle"},
		},
		{
			name:     "dot wildcard",
			pattern:  "a.ple",
			input:    "apple\nample\naple\nbanana",
			expected: []string{"apple", "ample"},
		},
		{
			name:     "escaped characters",
			pattern:  "apple\\.com",
			input:    "apple.com\napplecom\napple.net",
			expected: []string{"apple.com"},
		},
		{
			name:     "word boundary",
			pattern:  "\\bapple\\b",
			input:    "apple\npineapple\napple pie\napples",
			expected: []string{"apple", "apple pie"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			result, err := Grep(tt.pattern, reader)
			if err != nil {
				t.Errorf("Grep() error = %v", err)
				return
			}
			if !compareSlices(result, tt.expected) {
				t.Errorf("Grep() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGrep_EdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		pattern     string
		input       string
		expectError bool
		expected    []string
	}{
		{
			name:        "empty pattern",
			pattern:     "",
			input:       "test input",
			expectError: true,
		},
		{
			name:        "invalid regex",
			pattern:     "[unclosed",
			input:       "test input",
			expectError: true,
		},
		{
			name:        "complex regex",
			pattern:     "^(apple|banana)\\s+.*$",
			input:       "apple pie\nbanana split\norange juice",
			expected:    []string{"apple pie", "banana split"},
		},
		{
			name:        "newlines in input",
			pattern:     "test",
			input:       "line1\ntest line\nline3\nanother test\nline5",
			expected:    []string{"test line", "another test"},
		},
		{
			name:        "special characters in plain text",
			pattern:     "test*",
			input:       "test*\ntest\n*test",
			expected:    []string{"test*", "test", "*test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			result, err := Grep(tt.pattern, reader)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("Grep() expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Grep() unexpected error = %v", err)
				return
			}
			
			if !compareSlices(result, tt.expected) {
				t.Errorf("Grep() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsRegexPattern(t *testing.T) {
	tests := []struct {
		pattern string
		isRegex bool
	}{
		{"apple", false},
		{"^apple", true},
		{"apple$", true},
		{"apple|banana", true},
		{"[Aa]pple", true},
		{"appl*e", true},
		{"a.ple", true},
		{"apple\\.com", true},
		{"\\bapple\\b", true},
		{"apple{2}", true},
		{"apple+", true},
		{"apple?", true},
		{"(apple)", true},
		{"apple", false},
		{"test", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.pattern, func(t *testing.T) {
			result := isRegexPattern(tt.pattern)
			if result != tt.isRegex {
				t.Errorf("isRegexPattern(%q) = %v, want %v", tt.pattern, result, tt.isRegex)
			}
		})
	}
}

// Helper function to compare slices
func compareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
} 