package main

import "fmt"

func calculator() {

	var choice string

	var left float64
	var right float64

	fmt.Println("first num: ")
	_, err := fmt.Scanln(&left)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println("which operation (+,-,*,/): ")
	_, err = fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println("second num: ")
	_, err = fmt.Scanln(&right)
	if err != nil {
		fmt.Println("error")
		return
	}

	switch choice {
	case "+":
		fmt.Println("result: ", left+right)
		break
	case "-":
		fmt.Println("result: ", left-right)
		break
	case "*":
		fmt.Println("result: ", left*right)
		break
	case "/":
		if right != 0 {
			fmt.Println("result: ", left/right)

		} else {
			fmt.Println("invalid")
		}
		break
	}
}
