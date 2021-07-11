package main

import (
	"fmt"
)

func main() {
	// var color map[string]string Another way of declaring a map
	colors := make(map[string]string)
	// colors := map[string]string{
	// 	"red":  "34bkjvb",
	// 	"blue": "askjfh3",
	// }
	colors["white"] = "fffff"
	delete(colors, "white")
	fmt.Println(colors)
}
