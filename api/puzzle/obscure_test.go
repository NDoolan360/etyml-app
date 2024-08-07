package main

import "testing"

func TestIsCompletelyObscured(t *testing.T) {
	obscurer := '_'
	tests := []struct {
		in       string
		expected bool
	}{
		{"", true}, {"a", false}, {"_", true},
		{"a_", false}, {"_a", false}, {"aa", false}, {"__", true},
	}

	for _, test := range tests {
		out := isCompletelyObscured(test.in, obscurer)
		if out != test.expected {
			t.Errorf("input: '%s', expected output '%t' but got '%t'", test.in, test.expected, out)
		}
	}
}

func TestIsCompletelyUnobscured(t *testing.T) {
	obscurer := '_'
	tests := []struct {
		in       string
		expected bool
	}{
		{"", true}, {"a", true}, {"_", false},
		{"a_", false}, {"_a", false}, {"aa", true}, {"__", false},
	}

	for _, test := range tests {
		out := isCompletelyUnobscured(test.in, obscurer)
		if out != test.expected {
			t.Errorf("input: '%s', expected output '%t' but got '%t'", test.in, test.expected, out)
		}
	}
}

func TestObscureStringAfterNth(t *testing.T) {
	obscurer := '_'
	tests := []struct {
		string   string
		pos      int
		expected string
	}{
		{"", 99, ""}, {"a", 0, "_"}, {"a", 1, "a"},
		{"aaa", 0, "___"}, {"aaa", 1, "a__"},
		{"*aa", 0, "*__"}, {"*aa", 1, "*a_"},
	}

	for _, test := range tests {
		out := obscureStringAfterNth(test.string, test.pos, obscurer)
		if out != test.expected {
			t.Errorf("input: '%s %d', expected output '%s' but got '%s'",
				test.string, test.pos, test.expected, out)
		}
	}
}
