package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func firstNonRepeatingCharacter(s string) rune {
	// Step 1: Create a frequency map
	frequency := make(map[rune]int)

	// Count occurrences of each character
	for _, char := range s {
		frequency[char]++
	}

	// Step 2: Find the first character that appears only once
	for _, char := range s {
		if frequency[char] == 1 {
			return char
		}
	}

	// Step 3: If no unique character found, return '_'
	return '_'
}

func main() {
	testCases := []string{
		"hello world",
		"hello",
		"beach",
	}

	for _, test := range testCases {
		fmt.Printf("\"%s\" reversed >> \"%s\"\n", test, reverseString(test))
	}

	testCasesNonRepeating := []string{
		"abacabad", // Output: 'c'
		"abcdef",   // Output: 'a'
		"aabbcc",   // Output: '_'
		"swiss",    // Output: 'w'
	}

	for _, test := range testCasesNonRepeating {
		fmt.Printf("In string \"%s\" non repeating >> \"%c\"\n", test, firstNonRepeatingCharacter(test))
	}

	fizzBuzz(100)
}
