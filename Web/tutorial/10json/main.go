package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	//从哪里解码，解码到哪里
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
}
func encodeHandler(w http.ResponseWriter, r *http.Request) {
	peter := User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
	}
	//编码到哪里，从哪里编码
	json.NewEncoder(w).Encode(peter)
}
func main() {

	http.HandleFunc("/decode", decodeHandler)
	http.HandleFunc("/encode", encodeHandler)

	http.ListenAndServe(":8080", nil)
}
