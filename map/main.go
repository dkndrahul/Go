package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#00000",
	// 	"green": "#4785847",
	// }
	// var  colors map[string]string // the other way of initializing the map

	colors := make(map[string]string) // The other way of cretaing maps
	colors["RED"] = "jkjahfa"

	fmt.Println(colors)
}
