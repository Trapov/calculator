package main

import "calculator/machine"
import "fmt"

func main() {

    fmt.Print("Enter your expression: ")

    var input string
    fmt.Scanln(&input)

    fmt.Println(machine.Process(input))

}
