package mostwantedletter

import "testing"

func TestCheckio(t *testing.T) {
	isTrue(t, checkio("Hello World!") == 'l')
	isTrue(t, checkio("How do you do?") == 'o')
	isTrue(t, checkio("One") == 'e')
	isTrue(t, checkio("Oops!") == 'o')
	isTrue(t, checkio("AAaooo!!!!") == 'a')
	isTrue(t, checkio("abe") == 'a')
}

func isTrue(t *testing.T, v bool) {
	if !v {
		t.Error("Expect true but false")
	}
}
