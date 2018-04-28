package main

import (
	"math/big"
	"fmt"
)

func myfactorial (x *big.Int) (temp *big.Int) {
temp = new(big.Int)
	if x.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}

		return temp.Mul(temp, myfactorial(x.Sub(x, big.NewInt(1))))

}


func main() {


	x := big.NewInt(2)

	//fmt.Println(x.Cmp(big.NewInt(4)))
	//fmt.Println(x.Mul(x,big.NewInt(2)))
	//fmt.Println(x.Sub(x,big.NewInt(1)))
	//y := fact(x)
	fmt.Println (myfactorial(x))
	fmt.Println("hello world")
}