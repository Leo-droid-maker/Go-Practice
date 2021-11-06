package main

import "fmt"

func main() {
	colors := map[string]string{ // - One way to declare a map
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// var colors map[string]string // - another way to declare

	// fmt.Println(colors["red"])

	// colors := make(map[int]string) // - the best way

	// colors[10] = "#ffffff"

	fmt.Println(colors["white"])
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for given color: %s is %s\n", color, hex)
	}
}
