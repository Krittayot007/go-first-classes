package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
}

func loginHandler(w http.ResponseWriter, req *http.Request) { // -> ที่คล้าย express  (res, req) 
	method := req.Method // built in http req
	if method == "POST" {
		jsonByte, err := io.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("invalid request"))
			return
		}
		fmt.Println("body", string(jsonByte))
		var user User
		err = json.Unmarshal(jsonByte, &user) // -> JSON.parse
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("invalid request"))
			return
		}
		fmt.Println("user", user.FirstName, user.Username)
		w.WriteHeader(200)
		w.Write(jsonByte)
	}

}

func main() {
	http.HandleFunc("/login", loginHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
