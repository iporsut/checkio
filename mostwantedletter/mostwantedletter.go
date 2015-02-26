package mostwantedletter

import "unicode"

func checkio(text string) rune {
	histogram := countLetter(text)
	return findMostLetter(histogram)
}

func countLetter(text string) [26]int {
	var histogram [26]int
	for _, r := range text {
		index := -1
		if unicode.IsLower(r) {
			index = int(r - 'a')
		} else if unicode.IsUpper(r) {
			index = int(r - 'A')
		}
		if index > -1 {
			histogram[index]++
		}
	}
	return histogram
}

func findMostLetter(histogram [26]int) rune {
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
