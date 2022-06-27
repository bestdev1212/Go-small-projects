package palind

import (
	"math"
	"unicode/utf8"
)

// result := GetP("aa")
func GetPalind(s string) string {
	stringLength := utf8.RuneCountInString(s)
	isPalindromeMatrix := make([][]int, stringLength)
	for i := range isPalindromeMatrix {
		isPalindromeMatrix[i] = make([]int, stringLength)
	}

	for i, outer := range isPalindromeMatrix {
		for j := range outer {
			if i == j {
				isPalindromeMatrix[i][j] = 1
			}

		}
	}

	palindromeLengthMatrix := make([][]int, stringLength)
	for i := range palindromeLengthMatrix {
		palindromeLengthMatrix[i] = make([]int, stringLength)
	}

	for i, outer := range palindromeLengthMatrix {
		for j := range outer {
			if i == j {
				palindromeLengthMatrix[i][j] = 1
			}

		}
	}

	for len := 2; len <= stringLength; len++ {
		for i := 0; i <= stringLength-len; i++ {
			j := i + len - 1

			if s[i] == s[j] {
				if len == 2 {
					isPalindromeMatrix[i][j] = 1
					palindromeLengthMatrix[i][j] = 2
				} else {
					if isPalindromeMatrix[i+1][j-1] == 1 {
						isPalindromeMatrix[i][j] = 1
						palindromeLengthMatrix[i][j] = 2 + palindromeLengthMatrix[i+1][j-1]
					} else {
						isPalindromeMatrix[i][j] = -1
						palindromeLengthMatrix[i][j] = int(math.Max(float64(palindromeLengthMatrix[i+1][j]), float64(palindromeLengthMatrix[i][j-1])))
					}

				}

			} else {
				isPalindromeMatrix[i][j] = -1
			}

		}
	}

	max_row_index := 0
	max_column_index := 0
	max := 0

	for i, outer := range palindromeLengthMatrix {
		for j := range outer {
			if palindromeLengthMatrix[i][j] > max && isPalindromeMatrix[i][j] == 1 {
				max = palindromeLengthMatrix[i][j]
				max_row_index = i
				max_column_index = j
			}
		}
	}

	return s[max_row_index : max_column_index+1]
}
