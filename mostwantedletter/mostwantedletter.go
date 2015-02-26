package mostwantedletter

import (
	"unicode"
	"unicode/utf8"
)

func checkio(text string) rune {
	var histogram [26]int
	for len(text) > 0 {
		index := -1
		r, size := utf8.DecodeRuneInString(text)
		if unicode.IsLower(r) {
			index = int(r - 'a')
		} else if unicode.IsUpper(r) {
			index = int(r - 'A')
		}
		if index > -1 {
			histogram[index]++
		}

		text = text[size:]
	}

	mostIndex := -1
	mostCount := -1
	for i, n := range histogram {
		if n > mostCount {
			mostIndex = i
			mostCount = n
		}
	}

	return rune(mostIndex) + 'a'
}
