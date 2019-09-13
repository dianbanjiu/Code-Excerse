package main

import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "welcome to my site\n")
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static", http.StripPrefix("/stati/", fs))
	http.ListenAndServe(":8080", nil)
}