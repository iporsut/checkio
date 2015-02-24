package housepassword

import "testing"

func isTrue(t *testing.T, v bool) {
	if !v {
		t.Error("Expect true but false")
	}
}
func isFalse(t *testing.T, v bool) {
	if v {
		t.Error("Expect false but true")
	}
}

func TestCheckio(t *testing.T) {
	isFalse(t, checkio("A1213pokl"))
	isTrue(t, checkio("bAse730onE"))
	isFalse(t, checkio("asasasasasasasaas"))
	isFalse(t, checkio("QWERTYqwerty"))
	isFalse(t, checkio("123456123456"))
	isTrue(t, checkio("QwErTy911poqqqq"))
}
