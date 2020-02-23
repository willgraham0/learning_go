package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface.
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow type.
type Cow struct{}

// Bird type.
type Bird struct{}

// Snake type.
type Snake struct{}

// Eat prints the food that the Cow eats.
func (c Cow) Eat() {
	fmt.Println("grass")
}

// Move prints the movement that the Cow makes.
func (c Cow) Move() {
	fmt.Println("walk")
}

// Speak prints the sound that the Cow makes.
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Eat prints the food that the Bird eats.
func (b Bird) Eat() {
	fmt.Println("worms")
}

// Move prints the movement that the Bird makes.
func (b Bird) Move() {
	fmt.Println("fly")
}

// Speak prints the sound that the Bird makes.
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Eat prints the food that the Snake eats.
func (s Snake) Eat() {
	fmt.Println("mice")
}

// Move prints the movement that the Snake makes.
func (s Snake) Move() {
	fmt.Println("slither")
}

// Speak prints the sound that the Snake makes.
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	// Create an animals hashmap.
	var animals map[string]Animal
	animals = make(map[string]Animal)

	for true {
		// Get user input.
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		commands := strings.Split(text, " ")

		command0 := commands[0]
		command1 := commands[1]
		command2 := strings.TrimSuffix(commands[2], "\n")

		if command0 == "newanimal" {
			var animal Animal
			if command2 == "cow" {
				animal = Cow{}
			} else if command2 == "bird" {
				animal = Bird{}
			} else if command2 == "snake" {
				animal = Snake{}
			} else {
				fmt.Println("Invalid animal.")
				continue
			}
			// Add the animal to the dictionary of animals.
			animals[command1] = animal
			fmt.Println("Created it!")

		} else if command0 == "query" {

			animal, found := animals[command1]
			if found == false {
				fmt.Println("That animal does not exist.")
				continue
			}
			if command2 == "eat" {
				animal.Eat()
			} else if command2 == "move" {
				animal.Move()
			} else if command2 == "speak" {
				animal.Speak()
			} else {
				fmt.Println("Invalid action.")
				continue
			}

		} else {
			fmt.Println("Invalid command.")
		}

	}
}
