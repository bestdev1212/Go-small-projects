package palind

import (
	"testing"
)

func TestPalind(t *testing.T) {
	inp := "saabatt"
	expected := "aba"
	result := GetPalind(inp)
	t.Log(result)
	if result != expected {
		t.Fatal("Error")
	}
}
