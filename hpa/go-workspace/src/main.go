package main

import (
	"fmt"
	"math"
	"net/http"
)

func sqrt() string  {
	x := 0.0001

	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)		
	}
	
	return "Code.education Rocks!" ;
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, sqrt())

}

func greeting(message string) string {
	return fmt.Sprintf("<b>%s</b>", message)
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Print("Server listening on port 8000/")
	http.ListenAndServe(":8000",nil)
	
	
}