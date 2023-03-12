package main

import "fmt"

func exampleTypeConversion() {
	var x int = 10
	var y float64 = 30.2
	// var z float64 = x + y // NB: this will not work because x and y are not the same type; we need to convert
	var z float64 = float64(x) + y
	//var d int = x + y // NB2: Same deal here; no automatic type promotion for numeric types in Go
	var d int = x + int(y)
	fmt.Println(z)
	fmt.Println(d)
}

func main() {
	fmt.Println("Hello, world")

	exampleTypeConversion()
}
