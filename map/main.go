package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#00000",
	// 	"green": "#4785847",
	// }
	// var  colors map[string]string // the other way of initializing the map

	// colors := make(map[string]string) // The other way of cretaing maps
	// colors["RED"] = "jkjahfa"

	colors := make(map[int]string)

	colors[10] = "fkajhdf"
	delete(colors, 10) // way to delete an element/ key

	fmt.Println(colors)
}
