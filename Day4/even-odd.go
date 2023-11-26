package main

import "fmt"

func main() {
	sum := 0
	for count := 1; count <= 100; count++ {
		if count%2 == 0 {
			sum = sum + count
		} else {
			sum = sum - count
		}
	}
	fmt.Println(sum)
}
