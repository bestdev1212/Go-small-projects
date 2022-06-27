package bigN

import (
	"strconv"
	"strings"
)

func GetBigN(inp int, v int) int {
	inp_ := strconv.Itoa(inp)
	// value := strconv.Itoa(v)
	digits := strings.Split(inp_, "")
	isPositive := digits[0] != "-"
	if !isPositive {
		digits = digits[1:]
	}
	var i int
	for i = 0; i < len(digits); i++ {
		digit, _ := strconv.Atoi(digits[i])
		if isPositive {
			if v > digit {
				break
			}
		} else {
			if v < digit {
				break
			}
		}
	}
	shiftArray(&digits, i, strconv.Itoa(v))
	if !isPositive {
		shiftArray(&digits, 0, "-")
	}
	ret, _ := strconv.Atoi(strings.Join(digits, ""))
	return ret
}

func shiftArray(array *[]string, position int, value string) {
	//  extend array by one
	*array = append(*array, "")

	// shift values
	copy((*array)[position+1:], (*array)[position:])

	// insert value
	(*array)[position] = value
}
