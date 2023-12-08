package main

import "fmt"

func Concat() {}

type User struct {
	Username string
}

func (u User) Concat2() {
	fmt.Println("username: ", u.Username)
}

func main() {
	var user User
	fmt.Println("Hello world")
	user.Concat2()
}
