package main

import "fmt"

func main ()  {

	//i := 1
	//for i <= 10 {
	//	fmt.Println(i)
	//	i += 1
	//}


	//for i:=1; i<=10 ; i++ {
	//	fmt.Println(i)
	//}


	//for i:=1; i<=10; i++ {
	//	if i % 2 == 0 {
	//		fmt.Println(i, "четное")
	//	} else {
	//		fmt.Println(i, "нечетное")
	//	}
	//}


	//for i:=0; i<=10 ; i++ {
	//
	//	switch i {
	//	case 0:
	//		fmt.Println("Zero")
	//	case 1:
	//		fmt.Println("One")
	//	case 2:
	//		fmt.Println("Two")
	//
	//	default:
	//		fmt.Println("Unknow number")
	//
	//	}
	//}

				//задача 1
	//i:=1
	//if i > 10 {
	//	fmt.Println("big")
	//} else {
	//	fmt.Println("small")
	//}

				//задача 2
	//for i:=1; i <= 100; i++ {
	//	if i % 3 == 0 {
	//		fmt.Println(i)
	//	}
	//}

				//задача 3
	for i:=1; i<=100; i++  {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
