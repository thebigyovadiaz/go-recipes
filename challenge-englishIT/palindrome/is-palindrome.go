package main

import (
	"fmt"
	"regexp"
	"strings"
)

// cleanStringRegexp is a function that clean punctuation, spaces, and convert to lowercase a string
func cleanString(s string) string {
	regex := regexp.MustCompile(`[^a-zA-Z]`)
	cleaned := regex.ReplaceAllString(s, "")

	// Convert to lowercase
	return strings.ToLower(cleaned)
}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func isPalindrome(s string) bool {
	cleaned := cleanString(s)
	reversed := reverseString(cleaned)
	return cleaned == reversed
}

func main() {
	testCases := []string{
		"A man, a plan, a-¿¡! canal;: Pan;a¡@ma",
		"race a car",
		"No lemon, no melon!",
	}

	for _, test := range testCases {
		fmt.Printf("Is \"%s\" a palindrome? %v\n", test, isPalindrome(test))
	}
}
