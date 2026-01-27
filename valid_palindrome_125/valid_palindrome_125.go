package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	if s == "" || s == " " || len(s) <= 1 {
		return true
	}

	var formattedString string
	for _, i := range s {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' || i >= '0' && i <= '9' {
			formattedString += strings.ToLower(string(i))
		}
	}

	if len(formattedString) <= 1 {
		return true
	}

	for index, value := range formattedString {
		var a = formattedString[(len(formattedString)-1)-index]
		if value != int32(a) {
			return false
		}
	}

	return true

	// best solution

	var newStr []rune
	for _, i := range s {
		if unicode.IsLetter(i) || unicode.IsDigit(i) {
			newStr = append(newStr, unicode.ToLower(i))
		}
	}
	l, r := 0, len(newStr)-1
	for l < r {
		if newStr[l] != newStr[r] {
			return false
		}
		l++
		r--
	}
	return true
}
