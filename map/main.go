package main

import "fmt"

type colorType map[string]string

func main() {
	colors := colorType{
		"red":    "#FF0000",
		"green":  "#4bf745",
		"white":  "#fff",
		"ablack": "#000",
	}

	delete(colors, "ablack")
	colors.print()
}

func (c colorType) print() {
	for color, hex := range c {
		fmt.Printf("Color is %s and hex is %s \n", color, hex)
	}
}
