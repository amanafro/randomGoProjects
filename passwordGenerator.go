package main

import "fmt"

func main() {
	PasswordGenerator()
}

func PasswordGenerator() {

	const (
		y = "iota"
		n
	)

	var input string

	var char = []byte("#abcdefghilkmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+={}[]|;:,.<>/")

	for index, char := range char {
		fmt.Printf("code: %d, char %c\n", index, char)
	}

	fmt.Println("what should it contain? ")
	fmt.Println("number? ")
	fmt.Scanln(&input)
	fmt.Println("capital alphabets? ")
	fmt.Scanln(&input)
	fmt.Println("small alphabets? ")
	fmt.Scanln(&input)
	fmt.Println("special character? ")
	fmt.Scanln(&input)
  
  switch input{
  case "y":
  }
}
