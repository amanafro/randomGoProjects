package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numBytes     = "0123456789"
	specialBytes = "!@#$%^&*()_+{}[]|;:,.<>?/~`"
)

var yesOrNo string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func PasswordGenerator(length int, useNum, useSpecial, useLetter bool) string {
	var allowedChars string
	if useLetter {
		allowedChars += letterBytes
	}
	if useNum {
		allowedChars += numBytes
	}
	if useSpecial {
		allowedChars += specialBytes
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	return string(b)
}

func main() {

	var length int

	fmt.Println("wanna password (y/N): ")
	fmt.Scanln(&yesOrNo)

	if yesOrNo == "y" {
		fmt.Println("the length? ")
		fmt.Scanln(&length)

		if length < 4 || length > 25 {
			fmt.Println("invalid, short or too long")
		} else {
			password := PasswordGenerator(length, true, true, true)
			fmt.Println("passowrd: ", password)
		}
	} else {
		fmt.Println("bye!!!")
	}
}
