package main

import (
	"fmt"
	"time"
)

// Increment a value by 1.
func increment(value *int) {
	*value++
}

// Negate a value.
func negate(value *int) {
	*value = -(*value)
}

// This main function has a race condition within it. There is an
// underlying object which is an integer of value 2. Two goroutines are
// executed which take the value at the address of the integer and
// modify it. One increments the integer by 1 and the other negates it
// (multiplies it by -1). The final value of the integer depends on which
// goroutine is executed first. This is the race condition.
// In the example below, the final value of `x` is either -3 or -1.
func main() {
	// Create an object.
	var x int = 2

	// Increment it.
	go increment(&x)

	// Negate it.
	go negate(&x)

	// Wait for both goroutines to finish executing.
	time.Sleep(1 * time.Second)

	// Print it.
	fmt.Println("x =", x)
}
