package main

import "fmt"

func main() {
	var x int = 10
	var y *int = &x

	fmt.Printf("%p = x\n%p = y value\n%p = y address \n", &x, y, &y)
}
