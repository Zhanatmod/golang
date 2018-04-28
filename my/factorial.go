package main

import (
	"fmt"
	"math/big"
)

/*import (
	"fmt"
	"math/big"
)

func main() {
	var f big.Int
	f.MulRange(1, 995)
	fmt.Println(&f)
}
*/

func main() {
	fmt.Println("Hello, playground")

	//n := big.NewInt(40)
	r := big.NewInt(995)

	fmt.Println(factorial(r))

}

func factorial(n *big.Int) (result *big.Int) {
	//fmt.Println("n = ", n)
	result = new(big.Int)

	switch n.Cmp(&big.Int{}) {
	case -1, 0:
		result.SetInt64(1)
	default:
		fmt.Println("n = ", n)
		result.Set(n)
		var one big.Int
		one.SetInt64(1)
		result.Mul(result, factorial(n.Sub(n, &one)))
	}
	return
}
