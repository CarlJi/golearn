package strs_test

import (
	"testing"

	"carlji.com/codereadings/strings"
)

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type FieldTest struct {
	s   string
	res []string
}

var FiledsFuncTest = []FieldTest{
	{"", []string{}},
	{"ABCDABCD", []string{"BCD", "BCD"}},
	{"BCDE", []string{"BCDE"}},
}

func TestFieldsFunc(t *testing.T) {
	preFunc := func(char rune) bool { return char == 'A' }
	for _, fft := range FiledsFuncTest {
		a := strs.FieldsFunc(fft.s, preFunc)
		if !eq(a, fft.res) {
			t.Errorf("FieldsFunc(%q) = %v, want %v", fft.s, a, fft.res)
		}
	}
}
