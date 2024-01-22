package main

import (
  "fmt"
  "math/rand"
) 

func main() {
  password := PasswordGenerator(12,true, true)
  fmt.Print("password generated: ", password)
}

const (
    letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
    numBytes = "0123456789"
)

var length int
var password []byte

func PasswordGenerator(length, num bool, special bool, letter bool) string {
  b := make([]byte, length)
  for i := range b {
    if letter {
      b[i] = letter[rand.Intn(len(letterBytes))]
    } else if special {
      b[i] = special[rand.Intn(len(specialBytes))]
    } else if num {
      b[i] = num[rand.Intn(len(numBytes))]
    }
  }
  return string(b)
}
