package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	var x int = 10
	y := 11

	// if
	if x > y {
		fmt.Printf("%d is greater than %d\n", x, y)
	} else if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("Both are equal\n")
	}

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

	for {
		fmt.Println(time.Now())
		if math.Mod(float64(time.Now().Second()), 2) == 0 {
			break
		}
		time.Sleep(time.Second * 1)
	}
}
