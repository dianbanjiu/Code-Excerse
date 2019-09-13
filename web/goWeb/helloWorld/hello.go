package main
import (
	"net/http"
	"fmt"
)
func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "congratulications, you've got requested %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}