package main

import "fmt"

func main() {
	colors := map[string]string{}
	//fmt.Println(colors["red"]) // ""
	colors["red"] = "#ff0000"
	//fmt.Println(colors["red"]) // #ff0000
	//delete(colors, "red")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}