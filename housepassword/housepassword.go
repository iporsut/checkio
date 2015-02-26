package housepassword

import "unicode"

func checkio(password string) bool {
	var foundUpper, foundLower, foundNumber bool

	if len(password) >= 10 {
		for _, r := range password {
			foundUpper = foundUpper || unicode.IsUpper(r)
			foundLower = foundLower || unicode.IsLower(r)
			foundNumber = foundNumber || unicode.IsNumber(r)
			if foundUpper && foundLower && foundNumber {
				return true
			}
		}
	}

	return false
}
