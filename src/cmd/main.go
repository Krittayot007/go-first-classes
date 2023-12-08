package main

import "fmt"

func main() {
	for i := 0; i < 9; i++ {

	}
	x := 0
	for x < 8 {
		fmt.Println("a")
	}
	// for {

	// } // infinite loop

	var arr [9]string
	for index, value := range arr {
		fmt.Println(index, value)
	}

}
