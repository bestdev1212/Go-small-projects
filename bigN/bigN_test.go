package bigN

import (
	"fmt"
	"testing"
)

func TestBigN(t *testing.T) {
	x := -13579
	n := 4
	expected := -134579
	result := GetBigN(x, n)
	fmt.Println(result)
	if result != expected {
		t.Fatal("Error")
	}
}
