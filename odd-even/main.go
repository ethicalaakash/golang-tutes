package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)
	for _, i := range numbers {
		if i%2 != 0 {
			fmt.Println(i, " is odd")
		} else {
			fmt.Println(i, " is even")
		}
	}
}
