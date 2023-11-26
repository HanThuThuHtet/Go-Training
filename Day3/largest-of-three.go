package main

import "fmt"

func main() {
	num1 := 1
	num2 := 2
	num3 := 0
	if num1 > num2 && num1 > num3 {
		fmt.Println(num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Println(num2)
	} else {
		fmt.Println(num3)
	}
}
