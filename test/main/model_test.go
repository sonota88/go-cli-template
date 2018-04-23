package model

import "testing"
import "../../src/main"

func Test_Add(t *testing.T){
	if model.Add(1, 2) != 3 {
		t.Error("1 + 2 should be 3")
	}
}
