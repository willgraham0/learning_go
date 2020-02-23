package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Name representation.
type Name struct {
	fname string
	lname string
}

func main() {
	names := make([]Name, 0)

	// Get the name of the file.
	var fileName string
	fmt.Println("Enter the name of the file:")
	fmt.Scan(&fileName)

	// Open file.
	f, _ := os.Open(fileName)

	// Read file line by line.
	scanner := bufio.NewScanner(f)
	var words []string
	for scanner.Scan() {
		words = strings.Fields(scanner.Text())
		n := Name{fname: words[0], lname: words[1]}
		names = append(names, n)
	}

	// Close file.
	f.Close()

	// Print the first and last names of all the names in the file.
	for _, name := range names {
		fmt.Println(name.fname, name.lname)
	}
}
