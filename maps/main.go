package main

import "fmt"

func main() {
	//var colors map[string]string
	colors := make(map[string]string)
	
	colors["black"] = "#000000"
	colors["white"] = "#ffffff"
	colors["yellow"] = "#ffff00"

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color, hex)
	}
}