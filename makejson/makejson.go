package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var person map[string]string
	person = make(map[string]string)

	fmt.Printf("Enter a name:\n")
	nameReader := bufio.NewReader(os.Stdin)
	name, _ := nameReader.ReadString('\n')

	fmt.Printf("Enter an address:\n")
	addressReader := bufio.NewReader(os.Stdin)
	address, _ := addressReader.ReadString('\n')

	person["name"] = name
	person["address"] = address
	jsonified, _ := json.Marshal(person)
	fmt.Println("Json representation:")
	fmt.Println(string(jsonified))
}
