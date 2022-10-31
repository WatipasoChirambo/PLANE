package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/ussd", func(w http.ResponseWriter, r *http.Request){
		if r.Method == "POST" {
			r.ParseForm()
			session_id := r.FormValue("sessionId")
			fmt.Fprintf(w, "Hello, %q",session_id)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}