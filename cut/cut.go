package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

const prefix = "Basic "

func main() {
	auth := "Basic R28xMTg6d2VsY29tZSBjdXQ="
	const prefix = "Basic "
	base64Decoded, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		fmt.Printf("error while base64 decoding: %v\n", err)
		return
	}

	decodedString := string(base64Decoded)
	fmt.Printf("decoded: %s\n", decodedString)
	// Output: Go118:welcome cut

	username, password, ok := strings.Cut(decodedString, ":")
	if !ok {
		fmt.Println("not a basic authentication format.")
		return
	}
	fmt.Printf("username: %q, password: %q\n", username, password)
	// username: "Go118", password: "welcome cut"
}
