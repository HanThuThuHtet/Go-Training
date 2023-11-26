package main

import "fmt"

func main() {
	username := "admin"
	password := "password123,"
	if username == "admin" && password == "password123," {
		fmt.Println("Login Successful")
	} else {
		fmt.Println("Login Failed")
	}
}
