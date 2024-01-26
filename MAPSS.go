package main

import "fmt"

func main() {
	menu := map[string]float64{"soup": 4.69,
		"pie":           5.99,
		"salad":         7.87,
		"toffee pading": 3.55}
	fmt.Println(menu)
	fmt.Println(menu["pie"])
}
