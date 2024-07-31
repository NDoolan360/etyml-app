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

// Check whether rune has diacritics.
func isDiacritic(r rune) bool {
	if r >= 0x0300 && r <= 0x036F {
		return true
	}
	switch r {
	case '\u00A8', // Diaeresis
		'\u00B4',                                                   // acute accent
		'\u00B8',                                                   // cedilla
		'\u00C0', '\u00C1', '\u00C2', '\u00C3', '\u00C4', '\u00C5', // Capital A..Z
		'\u00E0', '\u00E1', '\u00E2', '\u00E3', '\u00E4', '\u00E5', // small a..z
		'\u0100', '\u0101', // capital A with macron
		'\u012e', '\u012F': // small c with caron
		return true
	default:
		return false
	}
}

// Tells how many chars in common the best guess has
func largestGuess(term string, guesses []string) int {
	max := 0

	for _, guess := range guesses {
		term = normalize(term)
		guess = normalize(guess)

		termLength := len(term)
		guessLength := len(guess)
		count := 0

		termPos := 0
		guessPos := 0

		for termPos < termLength && guessPos < guessLength {
			if isDiacritic(rune(term[termPos])) {
				termPos++
				count++
				continue
			}
			if isDiacritic(rune(guess[guessPos])) {
				guessPos++
				continue
			}

			if term[termPos] == guess[guessPos] {
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

// Obscure string after nth character
func obscureStringAfterNth(s string, n int, obscurer rune) string {
	count := 0
	result := ""

	for _, r := range s {
		if count < n {
			count++
			result += string(r)
		} else {
			result += string(obscurer)
		}
	}

	return result
}

func isCompletelyObscured(s string, obscurer rune) bool {
	for _, r := range s {
		if r != obscurer {
			return false
		}
	}
	return true
}

func isCompletelyUnobscured(s string, obscurer rune) bool {
	for _, r := range s {
		if r == obscurer {
			return false
		}
	}
	return true
}

func sneakPeek(s string, obscurer rune) string {
	count := 1

	for _, r := range s {
		if !unicode.IsLetter(r) {
			count++
		} else {
			break
		}
	}

	return obscureStringAfterNth(s, count, obscurer)
}
