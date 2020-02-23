package main

import "fmt"

// BubbleSort the slice from smallest to largest.
func BubbleSort(sli []int) {
	length := len(sli)

	swapped := true
	for swapped == true {
		swapped = false

		for i := 0; i < length-1; i++ {
			value := sli[i]
			nextValue := sli[i+1]
			if value > nextValue {
				Swap(sli, i)
				swapped = true
			}
		}
	}
}

// Swap the contents of the slice at index `i` with the contents at
// index `i+1`.
func Swap(sli []int, index int) {
	value := sli[index]
	nextValue := sli[index+1]
	sli[index] = nextValue
	sli[index+1] = value
}

// GetSequence returns a slice from user-inputted data.
func GetSequence() []int {
	capacity := 10
	length := 11
	sli := make([]int, 0, capacity)

	for length > capacity {
		fmt.Println("Enter the size of your sequence (<=10):")
		fmt.Scanln(&length)
	}

	fmt.Println("Enter the sequence (one integer at a time):")
	var entry int
	for len(sli) < length {
		fmt.Scanln(&entry)
		sli = append(sli, entry)
	}
	return sli
}

func main() {
	// Enter sequence.
	sequence := GetSequence()

	// Sort the sequence.
	BubbleSort(sequence)

	// Print the sorted sequence.
	fmt.Println(sequence)
}
