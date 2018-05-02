package main

import "fmt"

func half() (int, bool)  {
	if x % 2 == 0 {
		return ( x, true)
	} else {
		return x/2, false
	}
	
}
func main()  {
	x := 4
	fmt.Println(half(x))
}
