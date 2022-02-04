package main

import "fmt"

type color map[string]string

func main() {
	//var colors map[string]string
	colors := color{
		"red":   "#ff0000",
		"green": "#ff00ff",
		"white": "#ffffff",
	}
	//colors := make(map[string]string)
	//colors["test"] = "test"
	//fmt.Println(len(colors))
	//delete(colors, "test")
	fmt.Println(len(colors))
	fmt.Println(colors)
	printMap(colors)
	colors.print()

}

func (c color) print() {
	for col, hex := range c {
		fmt.Print(col)
		fmt.Println(hex)
	}
}
func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Print(color)
		fmt.Println(hex)
	}
}
