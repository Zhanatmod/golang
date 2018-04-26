package main

import (
"fmt"
)

func main() {

	var x int
	x = 5
	y := fact(x)

	fmt.Println(y)
}

func fact(x int) int {

	if x == 1 {
		return 1
	}
	return x * fact(x - 1)

}
