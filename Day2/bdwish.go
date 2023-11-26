package main

import "fmt"

func main() {
	str1 := "Happy"
	age := 23
	str2 := "rd Birthday"
	fmt.Printf("%s %d%s", str1, age, str2)

	newStr := str1 + " " + fmt.Sprint(age) + str2
	fmt.Println("\n" + newStr)
}
