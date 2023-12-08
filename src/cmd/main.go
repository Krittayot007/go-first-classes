package main

import "fmt"

func main() {
	// Declaration
	x := 10
	if x >= 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// Declaration with statement
	if y := 9; y >= 0 {
		fmt.Println("true", y)
	} else {
		fmt.Println("false", y)
	}
	// out of scope
	// fmt.Println("y",y)

	// Declaration
	z := "hello"
	switch z {
	case "world":
		fmt.Println("world-z", z)
	case "hello":
		fmt.Println("hello-z", z)
		fallthrough
	case "test": // fallthrough
		fmt.Println("test-z", z)
	default:
		fmt.Println("default-z", z)
	}

	// Declaration with expresstion
	a := "hello"
	switch {
	case a == "world":
		fmt.Println("world-a", a)
	case a == "hello":
		fmt.Println("hello-a", a)
		fallthrough
	case a == "test": // fallthrough
		fmt.Println("test-a", a)
	default:
		fmt.Println("default-a", a)
	}
}
