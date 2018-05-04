package main

import "fmt"

//func add(args ...int) int {
//	max := args[0]
//	for i :=1; i < len(args); i++ {
//		if args[i] > max {
//		max = args[i]
//		}
//	}
//	return max
//}
//
//func main()  {
//	slice := []int{1,2,5,9}
//	fmt.Println(add(slice...))
//}


func add(args ...int) int {
	max := args[0]
	for _, v := range args {
		if v > max {
		max = v
		}
	}
	return max
}

func main()  {
	slice := []int{1,2,5}
	fmt.Println(add(slice...))
}