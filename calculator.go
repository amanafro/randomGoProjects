package main

import "fmt"

func main() {
	calculator()
}

func calculator() {

	var choice string

	fmt.Println("which operation (+,-,*,/): ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("error")
		return
	}

	switch choice {
	case "+":
		add()
		break
	case "-":
		minus()
		break
	case "*":
		multi()
		break
	case "/":
		div()
		break
	}
}

func add() float64 {
	var left float64
	var right float64

	fmt.Println("put the")

	return left + right
}

func minus() float64 {
	var left float64
	var right float64

	return left - right
}

func multi() float64 {
	var left float64
	var right float64

	return left * right
}

func div() float64 {

	var left float64
	var right float64

	if right != 0 {
		return left / right
	} else {
		fmt.Println("error")
	}

	return 0
}
