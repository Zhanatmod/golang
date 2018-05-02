package main

import "fmt"

func sum(xs[]int) int {
x := 0
	for _,v := range xs  {
		x += v
	}
	return x
}
func main()  {
	slice := []int{1,2,3,4}
	fmt.Println(sum(slice))
}