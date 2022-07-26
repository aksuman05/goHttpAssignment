package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)



func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the server!!")
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	user1 := User{}

	err := json.NewDecoder(r.Body).Decode(&user1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user1)
	err = json.NewEncoder(w).Encode(user1)
	if err != nil {
		fmt.Println(err)
	}

}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	user1 := User{
		Firstname: "Abhishek",
		Lastname: "Suman",
	}

	err := json.NewEncoder(w).Encode(user1)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "Application/json")

}

func main() {
	fmt.Println("Server is running")
	http.HandleFunc("/welcome", WelcomeHandler)
	http.HandleFunc("/post-user", PostHandler)
	http.HandleFunc("/get-user", GetHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}