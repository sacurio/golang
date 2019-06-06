package lib

import "unicode"

//Function that validate if the input string is a digit number
func IsNumber(y string) bool {
	for _, c := range y {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
