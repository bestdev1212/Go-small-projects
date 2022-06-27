package jsonEncode

import (
	"fmt"
	"testing"
)

func TestJsonEncode(t *testing.T) {
	inp := `{
	  "car": {
	    "color": "blue",
	    "year": 1999,
	    "broken": [
	      "left mirror",
	      "right door"
	    ]
	  }
	}`
	expected :=
		`[
	car/color="blue",
	car/year=1999,
	car/broken[0]="left mirror",
	car/broken[1]="right door"
]`
	MyEncoder := MyEncode{}
	MyEncoder.Encode(inp)
	result := MyEncoder.GetResult()
	fmt.Println(result)
	// fmt.Println(expected)
	if result != expected {
		t.Error("json encode  error")
	}
	// t.Log("jsong encode passed")
}
