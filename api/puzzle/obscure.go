package main

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

// Obscure string after nth character
func obscureStringAfterNth(s string, n int, obscurer rune) (result string) {
	count := 0

	for i, r := range s {
		if i < len(s)-1 && s[i+1] == '(' && s[len(s)-1] == ')' {
			// Cut short on upcoming translation
			// e.g. 2nd char on "mantra (मन्त्र)" -> "ma____"
			break
		} else if skipChar(r) {
			// Add to string
			result += string(r)
		} else if count < n {
			// Add to string
			result += string(r)
			count++
		} else {
			// Obscure
			result += string(obscurer)
		}
	}

	return result
}
