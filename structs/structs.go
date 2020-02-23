package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal is a representation of animals.
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat prints the food that the animal eats.
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move prints the movement that the animal makes.
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak prints the sound that the animal makes.
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	// Define the three animals.
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	// Create an animal hashmap
	var zoo map[string]Animal
	zoo = make(map[string]Animal)
	zoo["cow"] = cow
	zoo["bird"] = bird
	zoo["snake"] = snake

	fmt.Println("Enter the animal and the action")

	for true {
		// Get animal name and action from user.
		fmt.Println(">")
		reader := bufio.NewReader(os.Stdin)
		entry, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("There was a problem with what you entered.\n")
		} else {
			result := strings.Fields(entry)
			name := result[0]
			action := result[1]

			// Get the animal.
			animal := zoo[name]

			// Print the action.
			if action == "eat" {
				animal.Eat()
			} else if action == "move" {
				animal.Move()
			} else if action == "speak" {
				animal.Speak()
			} else {
				fmt.Println("That action does not exist.")
			}
		}
	}
}
