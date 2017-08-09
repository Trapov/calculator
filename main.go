package main

import "calculator/machine"
import (
	"fmt"
	"log"
)

func main() {

	fmt.Print("Enter your expression: ")

	var input string
	fmt.Scanln(&input)
	var answer, err = machine.Process(input)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(answer)
	}
}
