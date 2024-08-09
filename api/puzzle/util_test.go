package main

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"*bʰer-", "*bʰer-"}, {"barō", "baro"}, {"*barô", "*baro"}, {"baron", "baron"},
		{"ornāre", "ornare"}, {"मन्त्र", "मनतर"}, {"Nathan", "nathan"},
		{"àáâäǎãåāa", "aaaaaaaaa"}, {"èéêëěẽēėęe", "eeeeeeeeee"}, {"ķḻű", "klu"},
	}

	for _, test := range tests {
		out := normalize(test.in)
		if out != test.expected {
			t.Errorf("input: '%s', expected output '%s' but got '%s'", test.in, test.expected, out)
		}
	}
}

func TestSkipChar(t *testing.T) {
	tests := []struct {
		in       rune
		expected bool
	}{
		{'a', false}, {'d', false}, {'l', false}, {'n', false}, {'o', false},
		{'ß', false}, {' ', true}, {'*', true}, {'(', true}, {')', true},
	}

	for _, test := range tests {
		out := skipChar(test.in)
		if out != test.expected {
			t.Errorf("input: '%c', expected output '%t' but got '%t'", test.in, test.expected, out)
		}
	}
}

func TestLargestGuess(t *testing.T) {
	testGuesses := []string{"nail", "Door", "javascript", "Mantra"}
	tests := []struct {
		in       string
		expected int
	}{
		{"Nathan", 2}, {"nathan", 2},
		{"Doolan", 3}, {"doolan", 3},
		{"JAVASCRIPT", 10}, {"*javascript", 10},
		{"mantra (मन्त्र)", 6}, {"mantle", 4},
	}

	for _, test := range tests {
		out := largestGuess(test.in, testGuesses)
		if out != test.expected {
			t.Errorf("input: '%s', testGuesses:'%v', expected output '%d' but got '%d'",
				test.in, testGuesses, test.expected, out)
		}
	}
}
