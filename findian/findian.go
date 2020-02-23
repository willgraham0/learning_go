package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("Enter a string:\n")
	reader := bufio.NewReader(os.Stdin)
	entry, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("There was a problem with what you entered.\n")
	} else {
		trimmed := strings.TrimSpace(entry)
		lowerCase := strings.ToLower(trimmed)
		stringLength := len(lowerCase)

		if strings.HasPrefix(lowerCase, "i") &&
			lowerCase[stringLength-1:] == "n" &&
			strings.Contains(lowerCase, "a") {
			fmt.Printf("Found!\n")
		} else {
			fmt.Printf("Not Found!\n")
		}
	}
}
