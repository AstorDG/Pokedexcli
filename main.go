package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Hello, World!")
}

func cleanInput(input string) []string {
	lower_input := strings.ToLower(input)
	return strings.Fields(lower_input)
}
