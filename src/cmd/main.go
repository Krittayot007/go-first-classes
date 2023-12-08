package main

import (
	"fmt"
)

var b string = "code" // Package variable Internal
var B string = "code" // Package variable Exported
// c := "camp"
func main() {
	var x string
	var y string = "hello"
	var z = "world" // Infer Type
	a := "GO"       // Shorthand Declaration
	{
		// scope := "bloacking"
	}
	// a = 10 // Staticcally type language // assign not declare
	fmt.Println(x, y, z, a, b)
}

func ConcatString(a, b string) {
	// startTime := time.Now()

	test01 := "hello" + "world"
	// test11 := `hello ${cann't use}`

	// endTime := time.Now()
	// elapsedTime := endTime.Sub(startTime)
	// startTime2 := time.Now()
	test02 := fmt.Sprintf("%s %s", "arigato", "gopher")
	// endTime2 := time.Now()
	// elapsedTime2 := endTime2.Sub(startTime2)
	// fmt.Printf("test01-%s:%s \ntest02-%s:%s\n", test01, elapsedTime, test02, elapsedTime2)
	fmt.Println(test01, test02)

}
