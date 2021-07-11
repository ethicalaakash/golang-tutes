package main

import (
	"fmt"
)

func main() {
	// var color map[string]string Another way of declaring a map
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "34bkjvb",
		"blue":  "askjfh3",
		"white": "ffff",
	}

	printMap(colors)
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex coode for", color, "is", hex)
	}
}
