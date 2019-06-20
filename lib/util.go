package lib

import "unicode"

//Function that validate if the input string is a digit number
func IsNumber(y string) bool {
	for _, i := range y {
		if !unicode.IsDigit(i) {
			return false
		}
	}
	return true
}
