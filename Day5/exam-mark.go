package main

import "fmt"

func main() {
	classes := [][]int{}
	classA := []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10}
	classB := []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10}
	classC := []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10}
	classes = append(classes, classA, classB, classC)

	for i, class := range classes {
		for _, mark := range class {
			if mark == 100 {
				fmt.Printf("full mark , %d ,%d \n", mark, i)
			} else if mark >= 80 {
				fmt.Printf("flying color , %d ,%d \n", mark, i)
			} else if mark >= 40 {
				fmt.Printf("pass , %d ,%d \n", mark, i)
			} else {
				fmt.Printf("fail , %d ,%d \n", mark, i)
			}
		}
	}
}
