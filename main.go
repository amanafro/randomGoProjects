package main

import (
	"fmt"
)

func main() {
test:
	for val := 1; val < 7; val++ {

		var choice string

		fmt.Println("choice:")

		fmt.Printf(" 1.Invest \n 2.Passowrd \n 3.Purrency \n 4.Calculator \n")
		fmt.Println("X for exit")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			Invest()
		case "2":
			passowrd()
		case "3":
			currency()
		case "4":
			calculator()
		case "x":
			break test
		default:
			fmt.Println("Invalid")
		}
	}

}
