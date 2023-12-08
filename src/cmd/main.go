package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	str, res := Multiply(10, 20)
	fmt.Println(str, res)
}

func Concat(a, b string) string {
	return a + b
}

func Multiply(a int, b int) (string, int) {
	return "a x b", a * b
}
