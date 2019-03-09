package fizzbuzz_test

import (
	"github.com/hgsgtk/go-snippets/fizzbuzz"
	"testing"
)

func TestRun(t *testing.T) {
	tests := map[string]struct {
		num      int
		expected string
	}{
		"15で割り切れる場合FizzBuzz": {
			num:      45,
			expected: "FizzBuzz",
		},
		"5で割り切れる場合Buzz": {
			num:      40,
			expected: "Buzz",
		},
		"3で割り切れる場合Buzz": {
			num:      39,
			expected: "Fizz",
		},
		"15, 5, 3で割り切れない場合そのまま": {
			num:      37,
			expected: "37",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if actual := fizzbuzz.Run(tt.num); tt.expected != actual {
				t.Errorf("Run is expected '%s', but got '%s'", tt.expected, actual)
			}
		})
	}
}
