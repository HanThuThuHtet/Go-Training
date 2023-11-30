package main

import "fmt"

func main() {

	strSlice := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	fmt.Println(strSlice)
	strPointer(&strSlice[5])
	strPointer(&strSlice[7])
	strPointer(&strSlice[9])
	fmt.Println(strSlice)

}

func strPointer(str *string) {
	*str = "Pointer" + *str
}

// pointer string slice
// at least 10
// func declare with accept ref str slice parameter
// change value of indexes  5 , 7 , 9
// print out original str slice value in main fun
