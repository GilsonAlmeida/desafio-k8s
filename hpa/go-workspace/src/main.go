package main

import (
	"fmt"
	"math"
	"net/http"
)

func sqrt()  {
	x := 0.0001

	fmt.Sprintf("oi",x)
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
		
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	sqrt()
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