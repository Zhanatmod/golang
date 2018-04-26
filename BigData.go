package main



import (
"fmt"
	"math/big"
)

func main() {

	var x int
	x = 5
	y := fact(x)

	fmt.Println(y)
}

func fact (x int) big.Int {

	if x == 1 {
		return 1
	}
	return x * fact(x - 1)

}