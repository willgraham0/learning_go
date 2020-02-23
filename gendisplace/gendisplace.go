package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn returns a function that calculated distance using the
// unitial values of acceleration, velocity and displacement.
func GenDisplaceFn(acc, vel, dis float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*acc*math.Pow(t, 2) + vel*t + dis
	}
}

func main() {
	var acc, vel, dis, time float64
	fmt.Println("Enter acceleration:")
	fmt.Scanln(&acc)
	fmt.Println("Enter initial velocity:")
	fmt.Scanln(&vel)
	fmt.Println("Enter initial displacement:")
	fmt.Scanln(&dis)
	fmt.Println("Enter time:")
	fmt.Scanln(&time)

	fn := GenDisplaceFn(acc, vel, dis)

	fmt.Println("Displacement after time entered:")
	fmt.Println(fn(time))
}
