package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	name string
	age  int
)

func main() {
	var argsRaw = os.Args
	fmt.Printf("%v\n", argsRaw)

	// Define flags
	flag.StringVar(&name, "name", "John", "User's name")
	flag.IntVar(&age, "age", 30, "User's age")

	// Parse command-line arguments
	flag.Parse()

	// Access values of flags
	fmt.Printf("Name: %s \n", name)
	fmt.Printf("Age: %d \n", age)
}
