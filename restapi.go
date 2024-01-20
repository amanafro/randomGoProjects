package main

import (
  "fmt"
  "http"
)

func main() {
}

type User struct {
  ID int `json:"id"`
  Name string `json:"name"`
}

func get()  {
  http.HandleFunc("/user", getUser)  
  http.HandleFunc("/user", getUserByID)  
  http.HandleFunc("/create", createUser)  
  http.HandleFunc("/delete", deleteUser)  
  http.HandleFunc("/update", updateUser)  
  
  port := 3000
  fmt.Println("Server is running on: %d...\n", port)
  http.ListenAndServe(":"+ strcov.Itoa(port), nil) 
}


