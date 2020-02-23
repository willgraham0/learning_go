package main

import "fmt"

func main() {

	var entry float32
	fmt.Printf("Enter a floating point number:\n")

	_, err := fmt.Scan(&entry)
	if err != nil {
		fmt.Printf("You did not enter a floating point number.\n")
	} else {
		var truncated int
		truncated = int(entry)
		fmt.Printf("Truncated floating point number: \n%d\n", truncated)
	}
}
