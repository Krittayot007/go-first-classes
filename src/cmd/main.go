package main

import "fmt"

func main() {

	Lesson(0)
}

func Lesson(index int) {
	switch index {
	case 1:
		Case1()
	case 2:
		Case2()
	case 3:
		Case3()
	case 4:
		Case4()
	default:
	}
}

func Case1() {
	// var arr [4]bool
	// var slc []bool
	// fmt.Println(arr)
	// fmt.Println(slc)

	// var arr [4]bool
	// var slc []bool = []bool{false, false, false, false}
	// fmt.Printf("arr:%v len:%v cap:%v\n", arr, len(arr), cap(arr))
	// fmt.Printf("slc:%v len:%v cap:%v\n", slc, len(slc), cap(slc))

	// slc = append(slc, true)
	// arr = append(arr, true)
	// fmt.Printf("arr:%v len:%v cap:%v\n", arr, len(arr), cap(arr))
	// fmt.Printf("slc:%v len:%v cap:%v\n", slc, len(slc), cap(slc))
}
func Case2() {
	// capacity
	var slc1 [3]bool
	var slc2 = make([]bool, 3, 4)
	fmt.Printf("1: len:%v cap:%v slc1:%v \n", len(slc1), cap(slc1), slc1)
	fmt.Printf("2: len:%v cap:%v slc2:%v \n", len(slc2), cap(slc2), slc2)

	// slc2 = append(slc2, true, true)
	// fmt.Printf("2: len:%v cap:%v slc2:%v \n", len(slc2), cap(slc2), slc2)
	// slc2 = append(slc2, true)
	// fmt.Printf("2: len:%v cap:%v slc2:%v \n", len(slc2), cap(slc2), slc2)
	// slc2 = append(slc2, true)
	// fmt.Printf("2: len:%v cap:%v slc2:%v \n", len(slc2), cap(slc2), slc2)

}
func Case3() {
	// address
	var slc1 []bool
	var slc2 []bool
	fmt.Println(slc1)
	fmt.Println(slc2)

	slc3 := make([]bool, 4)
	fmt.Printf("3: len:%v cap:%v slc3:%v \n\n", len(slc3), cap(slc3), slc3)

	/* 	slc4 := slc3[0:2]
	   	fmt.Printf("4: len:%v cap:%v slc4:%v \n", len(slc4), cap(slc4), slc4)
	   	slc5 := slc3[0:3]
	   	fmt.Printf("5: len:%v cap:%v slc5:%v \n\n", len(slc5), cap(slc5), slc5) */

	/* 	slc4 := slc3[0:2]
	   	slc4 = append(slc4, true)
	   	fmt.Printf("4: len:%v cap:%v slc4:%v \n", len(slc4), cap(slc4), slc4)
	   	fmt.Printf("3: len:%v cap:%v slc3:%v \n", len(slc3), cap(slc3), slc3) */

	/* 	slc4 := slc3[0:2]
	   	slc4 = append(slc4, true, true)
	   	fmt.Printf("4: len:%v cap:%v slc4:%v \n", len(slc4), cap(slc4), slc4)
	   	fmt.Printf("3: len:%v cap:%v slc3:%v \n", len(slc3), cap(slc3), slc3) */

	/* 	slc4 := slc3[0:2]
	   	slc4 = append(slc4, true, true, false)
	   	fmt.Printf("4: len:%v cap:%v slc4:%v \n", len(slc4), cap(slc4), slc4)
	   	fmt.Printf("3: len:%v cap:%v slc3:%v \n", len(slc3), cap(slc3), slc3) */

	/* 	slc4 := slc3[0:2]
	   	slc4 = append(slc4, true, true)
	   	fmt.Printf("4: len:%v cap:%v slc4:%v  arr:%p [0]:%p [1]:%p [2]:%p [3]:%p \n", len(slc4), cap(slc4), slc4, slc4, &slc4[0], &slc4[1], &slc4[2], &slc4[3])
	   	fmt.Printf("3: len:%v cap:%v slc3:%v  arr:%p [0]:%p [1]:%p [2]:%p [3]:%p \n", len(slc3), cap(slc3), slc3, slc3, &slc3[0], &slc3[1], &slc3[2], &slc3[3])
	*/

	/*
		 	slc4 := slc3[0:2]
			slc4 = append(slc4, true, true, false)
			fmt.Printf("4: len:%v cap:%v slc4:%v  arr:%p [0]:%p [1]:%p [2]:%p [3]:%p \n", len(slc4), cap(slc4), slc4, slc4, &slc4[0], &slc4[1], &slc4[2], &slc4[3])
			fmt.Printf("3: len:%v cap:%v slc3:%v      arr:%p [0]:%p [1]:%p [2]:%p [3]:%p \n", len(slc3), cap(slc3), slc3, slc3, &slc3[0], &slc3[1], &slc3[2], &slc3[3])
	*/
}

func Case4() {
	// assign value
	var slc1 []bool
	var slc2 = make([]bool, 4)
	fmt.Printf("1: len:%v cap:%v slc1:%v \n", len(slc1), cap(slc1), slc1)
	fmt.Printf("2: len:%v cap:%v slc2:%v \n", len(slc2), cap(slc2), slc2)

	// slc1[0] = true
	// slc1[1] = true
	// fmt.Printf("1: len:%v cap:%v slc1:%v \n", len(slc1), cap(slc1), slc1)
	// fmt.Printf("2: len:%v cap:%v slc2:%v \n", len(slc2), cap(slc2), slc2)

}
