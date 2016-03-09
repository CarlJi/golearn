package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "ABCDEFGABCDEFGF"
	res := strings.FieldsFunc(s, isSplict)

	fmt.Printf("Result: %v ", res)
}

func isSplict(char rune) bool {
	switch char {
	case 'B', 'F':
		return true
	}
	return false
}
