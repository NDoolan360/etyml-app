package main

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func normalize(s string) string {
	result, _, _ := transform.String(
		transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC),
		s,
	)
	return strings.ToLower(result)
}

func skipChar(r rune) bool {
	return !unicode.IsLetter(r)
}

// Tells how many chars in common the best guess has
func largestGuess(term string, guesses []string) (max int) {
	for _, guess := range guesses {
		term = normalize(term)
		guess = normalize(guess)

		termLength := len(term)
		guessLength := len(guess)
		count := 0

		termPos := 0
		guessPos := 0

		for termPos < termLength && guessPos < guessLength {
			if skipChar(rune(term[termPos])) {
				termPos++
			} else if skipChar(rune(guess[guessPos])) {
				guessPos++
			} else if term[termPos] == guess[guessPos] {
				count++
				termPos++
				guessPos++
			} else {
				break
			}
		}

		if count > max {
			max = count
		}
	}

	return max
}
