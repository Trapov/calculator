package main

import "calculator/translator"
import "fmt"

func main(){
	

	fmt.Print("Enter your expression: ")

    var input string
    fmt.Scanln(&input)

	fmt.Println(translator.Process(input))

}