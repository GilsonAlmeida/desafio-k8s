package main

import ("testing" 
		"fmt")


func TestSqrt(t *testing.T) {
	
	
	result := sqrt()

	fmt.Print("teste")
	if result != "Code.education Rocks!" {
		t.Errorf("expected Code.education Rocks! but got %s", result)
	}

}