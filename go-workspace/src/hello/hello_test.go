package main

import "testing"


func TestGreeting(t *testing.T) {
	result := greeting("Code.education Rocks!")
	if result != "<b>Code.education Rocks!</b>" {
		t.Errorf("expected <b>Code.education Rocks!</b>! but got %s", result)
	}

}