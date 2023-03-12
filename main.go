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

func arraysExample() {
	var x [3]int // array of 3 ints which are all initialized to zero value
	fmt.Println(x)

	var x2 = [3]int{10, 20, 30} // declare explicit values for the arrays
	fmt.Println(x2)

	var x3 = [12]int{1, 5: 4, 6, 10: 100, 15} // sparsely populated array
	fmt.Println(x3)

	var x4 = [...]int{11, 13, 15} // we don't have to specify the size of the array when initializing with a literal
	fmt.Println(x4)

	fmt.Println(x4 == x2) // compare with == or !=

	fmt.Println(len(x3)) // get the length
}

func slicesExample() {
	var x = []int{10, 20, 30} // almost the same as an array, just leave out the length
	fmt.Println(x)

	var y = []int{11, 21, 32}
	//fmt.Println(x == y) // NB: it's impossible to compare slices to each other...
	fmt.Println(y == nil) // 	... but we can compare them to nil

	x = append(x, 11, 20) // grow the slice
	fmt.Println(x)
}

func main() {
	fmt.Println("Hello, world")

	exampleTypeConversion()

	arraysExample()
}
