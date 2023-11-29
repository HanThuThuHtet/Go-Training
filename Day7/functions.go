package main

import (
	"fmt"
	"math"
)

func main() {
	greeting := greet("Han")
	fmt.Println(greeting)

	square := calculateSquare(4)
	fmt.Println(square)

	sum, diff := sumAndDifference(2, 3)
	fmt.Println(sum, diff)

	quotient, remainder := dividedWithRemainder(8, 3)
	fmt.Println(quotient, remainder)
}

func greet(name string) string {
	return "Hello " + name
}

func calculateSquare(num int) int {
	return int(math.Pow(float64(num), 2))
}

func sumAndDifference(a int, b int) (int, int) {
	return a + b, a - b
}

func dividedWithRemainder(a int, b int) (int, int) {
	return a / b, a % b
}
