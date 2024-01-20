package main

import (
	"fmt"
	"math"
)

func main() {
	invest()
}

func invest() {
	var investamount float64
	var rate float64
	var chosenTime float64

	fmt.Println("How much would you like to invest? ")
	_, err := fmt.Scanln(&investamount)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println("What is the investment rate? ")
	_, err = fmt.Scanln(&rate)
	if err != nil {
		fmt.Println("Error reading the investment rate:", err)
		return
	}
	fmt.Println("What is the choose time(months)? ")
	_, err = fmt.Scanln(&chosenTime)
	if err != nil {
		fmt.Println("Error reading the chosen time:", err)
		return
	}

	result := investamount * math.Pow(1+rate/100, chosenTime)

	fmt.Println("The return is: ")
	fmt.Println(math.RoundToEven(result))
}
