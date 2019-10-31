package main

import (
	"fmt"
)

func main() {
	// print information to console
	fmt.Println("Hello world")

	// variables
	var v int
	v = 1
	fmt.Println("v = ", v)

	var x int = 10
	fmt.Println("x = ", x)
	y := 11
	fmt.Println("y =", y)
	xy := x * y
	fmt.Println("x * y =", xy)

	// if
	if x > y {
		fmt.Printf("%d is greater than %d\n", x, y)
	} else if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("Both are equal\n")
	}

	// Arrays
	// Go's arrays are values. An array variable denotes the entire array; it is not a pointer to the first array element (as would be the case in C)
	var a [5]int
	fmt.Println(a)

	a[2] = 7
	fmt.Println(a)

	//strs := [...]string{"Penn", "Teller"} //the compiler count the array elements
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// Slices
	c := []int{1, 2, 3, 5, 5}
	fmt.Println(c)
	c = append(c, 6)
	fmt.Println(c)

	d := c[2:4]
	fmt.Println(d)

	// map
	edges := make(map[string]int)
	edges["triangle"] = 3
	edges["square"] = 4
	fmt.Println(edges)
	fmt.Println(edges["triangle"])

	delete(edges, "triangle")
	fmt.Println(edges)

	// loops
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Print("\n")

	j := 0
	for j < 5 {
		fmt.Print(j)
		j++
	}
	fmt.Print("\n")

	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println("Index:", index, "Value:", value)
	}
}
