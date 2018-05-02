package main

import "fmt"

//func main()  {
//	xs := []float64{98, 93, 77, 82, 83}
//	total := 0.0
//	for _, i := range xs {
//		total += i
//	}
//	fmt.Println(total / float64(len(xs)))
//}


//	func average(xs[]float64)float64  {
//		total := 0.0
//		for _,v:=range xs {
//			total += v
//		}
//		return total / float64(len(xs))
//	}
//func main()  {
//	someOtherName := []float64{98,93,77,82,83}
//	fmt.Println(average(someOtherName))
//}


//func f(x int)  {
//	fmt.Println(x)
//}
//
//func main()  {
//	x := 6
//	f(x)
//}

//func main()  {
//	fmt.Println(f1())
//}
//func f1() int {
//	return f2()
//}
//func f2() (r int) {
//	r = 1
//	return
//}

					//Возврат нескольких значений
//func f() (int, int) {
//	return 5, 6
//}
//func main()  {
//	x, y := f()
//	fmt.Println(x, y)
//}

					//Переменное число аргументов функции
//func add(args ...int) int  {
//	total := 0
//	for _, v := range args {
//		total += v
//	}
//	return total
//}
//func main()  {
//	fmt.Println(add(1,2,3))
//}

					//Передача среза int-ов в функцию
//func add(args ...int) int  {
//	total := 0
//	for _, v := range args {
//		total += v
//	}
//	return total
//}
//func main()  {
//	xs := []int{1,2,3}
//	fmt.Println(add(xs...))
//}

					//Замыкания
//func main()  {
//	add := func(x, y int) int {
//		return x + y
//	}
//	fmt.Println(add(1, 1))
//}


//func main()  {
//	x := 0
//	increment := func() int {
//		x++
//		return x
//	}
//	fmt.Println(increment())
//	fmt.Println(increment())
//	fmt.Println(increment())
//}

//func makeEvenGenerator() func() uint {
//	i := uint(0)
//	return func() (ret uint) {
//		ret = i
//		i += 2
//		return
//	}
//}
//func main()  {
//	nextGen := makeEvenGenerator()
//	fmt.Println(nextGen())
//	fmt.Println(nextGen())
//	fmt.Println(nextGen())
//}

						//Рекурсия
//func factorial(x uint) uint {
//	if x == 0 {
//		return 1
//	}
//return x * factorial(x-1)
//}
//func main()  {
//	fmt.Println(factorial(5))
//}

						//Отложенный вызов
//func first()  {
//	fmt.Println("1st")
//}
//func second()  {
//	fmt.Println("2nd")
//}
//func main()  {
//	defer second()
//	first()
//}

						//Паника и восстановление
func main()  {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}