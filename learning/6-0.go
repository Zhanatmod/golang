package main

import "fmt"

func main() {
	//	var x [5]int
	//	x[4] = 100
	//	fmt.Println(x[4])

	//	var x [5]float64
	//	x[0] = 98
	//	x[1] = 93
	//	x[2] = 77
	//	x[3] = 82
	//	x[4] = 83
	//	var total float64 = 0
	//	for i:=0; i<len(x); i++ {
	//		total += x[i]
	//	}
	//fmt.Println(total / float64(len(x)))

	//x := [5]float64{98, 93, 77, 82, 83}
	//var total float64 = 0
	//for _, value := range x {
	//	total += value
	//}
	//fmt.Println(total / float64(len(x)))

	//var x []float64
	//x := make([]float64, 5)

	//x := make([]float64, 5, 10)

	//arr := [5]float64{1,2,3,4,5}
	//x:=arr[0:5]

	//slice1 := []int{1,2,3}
	//slice2 := append(slice1, 4, 5)
	//fmt.Println(slice1, slice2)

	//slice1 := []int{1,2,3,4}
	//slice2 := make([]int, 2)
	//copy(slice2, slice1)
	//fmt.Println(slice1, slice2)
	//x := make(map[string]int)
	////var x map[string] int
	//x["key"] = 10
	//fmt.Println(x["key"])

//x := make(map[int]int)
//x[2] = 11
//x[3] = 13
//delete(x, 3)
//fmt.Println(x[3])

//element := make(map[string]string)
//element["Hi"] = "Hilarry"
//element["DD"] = "Daddy"
//	if name, ok := element["Hi"]; ok {
//		fmt.Println(name, ok)
//	}

//element := map[string]string{
//	"H":"Hillary",
//	"Z":"Zhopa",
//	}


	element := map[string]map[string]string{
		"H":map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He":map[string]string{
			"name":"Helium",
			"state":"gas",
		},
		"Zh":map[string]string{
			"name":"Zhan",
			"state":"solid",
		},
	}
	if el, ok := element["He"]; ok {
		fmt.Println(el["name"], el["state"])
	}

}



