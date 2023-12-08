package main

import "fmt"

type Event struct {
	E string
}

func main() {
	var obj1 map[string]interface{}
	fmt.Println("obj1", obj1)

	obj2 := make(map[string]interface{})
	fmt.Println("obj2", obj2)

	// obj1["value"] = 1
	// obj2["value"] = 1

	obj3 := make(map[int]interface{})
	obj4 := make(map[bool]interface{})
	obj5 := make(map[string]int)
	obj6 := make(map[string]string)
	obj7 := make(map[string]Event)

	for key, value := range obj2 {
		fmt.Println("key", key, "value", value)
	}

	fmt.Println(obj1, obj2, obj3, obj4, obj5, obj6, obj7)

}
