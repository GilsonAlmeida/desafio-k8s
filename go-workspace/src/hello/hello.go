package main

import (
	"fmt"	
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func greeting(message string) string {
	return fmt.Sprintf("<b>%s</b>", message)
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Print("Server listening on port 8000/")
	http.ListenAndServe(":8000",nil)
	
	
}