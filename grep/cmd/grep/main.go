package main

import (
	"flag"
	"fmt"
	"os"

	"grep"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <pattern> [file]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Search for lines containing the pattern in stdin or file\n")
		fmt.Fprintf(os.Stderr, "Patterns with regex metacharacters (., *, +, ?, ^, $, [, ], (, ), |, \\, {, }) are treated as regex\n\n")
		fmt.Fprintf(os.Stderr, "Examples:\n")
		fmt.Fprintf(os.Stderr, "  cat file.txt | %s apple          # Plain text search from stdin\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s apple file.txt                # Plain text search in file\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  cat file.txt | %s \"^apple\"       # Regex search from stdin\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s \"apple|banana\" file.txt      # Regex search in file\n", os.Args[0])
	}

	flag.Parse()

	// Check if pattern is provided
	if flag.NArg() < 1 || flag.NArg() > 2 {
		fmt.Fprintf(os.Stderr, "Error: pattern is required, file is optional\n")
		flag.Usage()
		os.Exit(1)
	}

	pattern := flag.Arg(0)
	var input *os.File
	var err error

	// Determine input source
	if flag.NArg() == 1 {
		// Read from stdin
		input = os.Stdin
	} else {
		// Read from file
		filePath := flag.Arg(1)
		input, err = os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file '%s': %v\n", filePath, err)
			os.Exit(1)
		}
		defer input.Close()
	}

	// Search for pattern
	matches, err := grep.Grep(pattern, input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Output matching lines
	for _, line := range matches {
		fmt.Println(line)
	}
} 