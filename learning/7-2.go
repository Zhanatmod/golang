package main

import "fmt"

func half(x int) (int, bool)  {

	if x % 2 == 0 {
		return 1, true
	} else {
		return 0, false
	}

}
func main()  {
fmt.Println("Введите число: ")
var input int
fmt.Scanf("%d", &input)
	fmt.Println(half(input))
}

