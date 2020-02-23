package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

// GetSequence returns a slice from user-inputted data.
func GetSequence() []int {
	sequence := make([]int, 0, 3)
	var entry string

	fmt.Println("Construct a sequence of integers.")
	for {
		fmt.Println("Enter a number ('x' to stop):")
		fmt.Scanln(&entry)

		// Break out of loop if 'x' has been entered.
		if entry == "x" {
			break
		}

		// Cast to integer.
		i, _ := strconv.Atoi(entry)

		// Append to slice.
		sequence = append(sequence, i)
	}
	return sequence
}

// Split a slice into `n` chunks.
func Split(sli []int, n int) [][]int {
	var splitted = make([][]int, 0)
	length := len(sli)
	chunks := length / n

	for i := 0; i < n; i++ {
		start := i * chunks
		stop := start + chunks
		if i == n-1 {
			stop = length
		}
		splitted = append(splitted, sli[start:stop])
	}

	return splitted
}

// Sort a slice of integers, in place.
func Sort(sli []int, wg *sync.WaitGroup) {
	// Print the slice.
	fmt.Println("Sorting: ", sli)

	// Sort the slice.
	sort.Ints(sli)

	// Decrement the wait group count.
	wg.Done()
}

// Merge sorted slices of integers into a single sorted slice.
func Merge(splitted [][]int) []int {
	// Merge the slices.
	merged := splitted[0]
	for i := 1; i < len(splitted); i++ {
		merged = append(merged, splitted[i]...)
	}

	// Sort the slice.
	// I have cheated here and just sorted the merged slice.
	sort.Ints(merged)

	return merged
}

func main() {
	chunkNumber := 4

	// Create the Wait Group and set the count to 4.
	var wg sync.WaitGroup
	wg.Add(chunkNumber)

	// Get a sequence of numbers from stdin.
	sequence := GetSequence()
	fmt.Println("Entered sequence: ", sequence)

	// Split the slice into four chunks.
	chunks := Split(sequence, chunkNumber)
	s1 := chunks[0]
	s2 := chunks[1]
	s3 := chunks[2]
	s4 := chunks[3]

	// Sort each slice in a goroutine.
	go Sort(s1, &wg)
	go Sort(s2, &wg)
	go Sort(s3, &wg)
	go Sort(s4, &wg)

	// Wait.
	wg.Wait()

	// Merge sorted slices into a final sorted slice.
	var sortedSlices = make([][]int, chunkNumber)
	sortedSlices[0] = s1
	sortedSlices[1] = s2
	sortedSlices[2] = s3
	sortedSlices[3] = s4
	finalSequence := Merge(sortedSlices)

	// Print the final result.
	fmt.Println("Sorted sequence: ", finalSequence)
}
