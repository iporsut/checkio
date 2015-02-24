package housepassword

import (
	"unicode"
	"unicode/utf8"
)

func checkio(password string) bool {
	var foundUpper, foundLower, foundNumber bool

	if len(password) >= 10 {
		for len(password) > 0 {
			r, size := utf8.DecodeRuneInString(password)
			foundUpper = foundUpper || unicode.IsUpper(r)
			foundLower = foundLower || unicode.IsLower(r)
			foundNumber = foundNumber || unicode.IsNumber(r)
			if foundUpper && foundLower && foundNumber {
				return true
			}
			password = password[size:]
		}
	}

	return false
}
