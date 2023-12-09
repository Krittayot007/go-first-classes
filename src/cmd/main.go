package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var boo bool
	var str string
	var x int
	var y int32 // 2^8 == 256
	var xx uint8
	var by byte // alias for uint8
	hello, _ := json.Marshal("hello")
	fmt.Println("hello", hello)
	var z int16 // 2^16
	var a int32 // 2^32
	var rn rune // alias for int32
	rn = '🥰'
	fmt.Println("rune", rn)
	fmt.Printf("rune %c\n", rune(129392))
	var b int64 // 2^64
	var f32 float32 = 0.7999999
	var f64 float64 = 0.888888

	fmt.Printf("boo:%T str:%T x:%T y:%T z:%T a:%T b:%T xx:%T by:%T  f32:%.2f f64%.4f\n", boo, str, x, y, z, a, b, xx, by, f32, f64)
}

func WorkShop() {
	/*
		1. Create Variable type of int and int8 and try to add it to gether
		2. Create Variable type of int and int32 and try to add it to gether
		3. Create Variable type of rune asign Emoji within and Print it
		4. Take result of 3. and try Print Emoji to console
		5. Add this expression to your code and try to work then explain it ====> fmt.Printf("Value : %v\n Type : %T %b %d %o %O %p \n", v,v,v,v,v,v,&v)
	*/
}
