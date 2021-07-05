package main

import "fmt"

func is_even(arr []int) {
	for _, v := range arr {
		if v%2 == 0 {
			fmt.Println(v, " is even")
		} else {
			fmt.Println(v, " is odd")
		}
	}
}
func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	is_even(numbers)
}
