package main

import "fmt"

func main() {

	//  *type -> pointer

	//  *variable -> value address
	//  &variable -> address

	var x int = 10
	var y *int = &x
	// var z int = *y -> ตัวนี้คืออะไร -> กลับมาเป็น value ของ pointer นั้นๆ

	fmt.Printf("%p = x\n%p = y value\n%p = y address \n", &x, y, &y)
}
