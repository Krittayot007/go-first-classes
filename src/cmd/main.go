package main

import "fmt"

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	Age       int    `json:"age"`
}

func main() {
	var user User
	// Dot notation
	fmt.Println(user.Age)

	user2 := User{Username: "user2", Password: "1234"}
	fmt.Println(user2.Username)

}
