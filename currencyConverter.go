package main

import "fmt"

func currency() {

	var amount float64

	var chf_usd float64 = 1.16
	var chf_eur float64 = 1.06
	var chf_gbp float64 = 0.91
	var choice string

	fmt.Println("Which currey would you like to convert?")
	fmt.Println("1. CHF => EUR")
	fmt.Println("2. CHF => USD")
	fmt.Println("3. CHF => GBP")

	_, err := fmt.Scanln(&choice)
	if err != nil {
		return
	}

	fmt.Println("How much? ", choice)

	_, err = fmt.Scanln(&amount)
	if err != nil {
		return
	}

	switch choice {

	case "1":
		result := amount * chf_eur
		fmt.Print("The converted amount is: €", result)
		break

	case "2":
		result := amount * chf_usd
		fmt.Print("The converted amount is: $", result)
		break

	case "3":
		result := amount * chf_gbp
		fmt.Print("The converted amount is: £", result)
		break
	}
}
