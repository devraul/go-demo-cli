package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("\n--------Args-------")
	for _, arg := range os.Args {
		fmt.Println("- " + arg)
	}
	fmt.Println("----end of args----\n")
}