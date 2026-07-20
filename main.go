package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Lowercase and split a string on white space
func cleanInput(input string) []string {
	lower_input := strings.ToLower(input)
	return strings.Fields(lower_input)
}

func main() {
	user_input_scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		user_input_scanner.Scan()
		user_input := user_input_scanner.Text()
		split_user_input := cleanInput(user_input)
		fmt.Printf("Your command was: %v\n", split_user_input[0])
	}
}
