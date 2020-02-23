package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	// Variable slice initialized with a length of 0 and capacity of 3.
	// Instructions say length of 3 but that doesn't make much sense.
	sli := make([]int, 0, 3)

	var entry string

	// Create a "while" loop.
	for {
		// Get user input.
		fmt.Printf("Enter a number ('X' to stop):\n")
		fmt.Scanln(&entry)

		// Break out of loop if 'X' has been entered.
		if entry == "X" {
			break
		}

		// Cast to integer.
		i, _ := strconv.Atoi(entry)
		// Append to slice.
		sli = append(sli, i)
		// Sort slice in ascending order and print.
		sort.Ints(sli)
		fmt.Println(sli)
	}
}
