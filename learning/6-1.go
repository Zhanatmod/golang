package main

import "fmt"

func main()  {

//	x := []int{
//		48,96,86,68,
//		57,82,63,70,
//		37,34,83,27,
//		19,97,9,17,
//	}
//min := x[0]
//	for i:=1; i < len(x); i++ {
//		if  x[i] < min {
//			min = x[i]
//		}
//	}
//	fmt.Println(min)


//		x := []int{
//			48,96,86,68,
//			57,82,63,70,
//			37,34,83,27,
//			19,97,9,17,
//		}
//		min := []int{x[0]}
//	for i, value := range x {
//		if value < min[0] {
//			min[0] = x[i]
//		}
//	}
//fmt.Println(min)



	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97,9,17,
	}
	min := x[0]
	for i, value := range x {
		if value < min {
			min = x[i]
		}
	}

	fmt.Println(min)
}