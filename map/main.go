package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "4hs392",
	}

	// var colors map[string]string

	// colors := make(map[string] string)

	// structName.white
	// colors["white"] = "#ffffff"

	// delete(colors, "white")

	printMap(colors)
}

func printMap (color map[string]string) {
	// for value, key := range color{
	for color, hex := range color {
		fmt.Println("hex code for", color, "is", hex)
	}
}