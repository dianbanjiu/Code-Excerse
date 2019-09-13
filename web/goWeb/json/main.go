package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type User struct{
	Firstname string `json:firestname`
	Lastname string `json:lastname`
	Age int `json: age`
}

func main(){
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request){
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request){
		peter := User{
			Firstname : "John",
			Lastname : "Doe",
			Age : 28,
		}
		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":8080", nil)
}