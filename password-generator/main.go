package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command line flags
	length := flag.Int("length", 8, "Password length (8-100)")
	includeSymbols := flag.Bool("include-symbols", false, "Include symbols in password")
	
	// Parse command line arguments (help is provided by default)
	flag.Parse()
	
	if err := ValidateLength(*length); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
	// Generate and output a random password
	password := GeneratePassword(*length, *includeSymbols)
	fmt.Println(password)
} 
