package main

import (
	"bytes"
	"context"
	"reflect"
	"testing"
)

var testNode = Node{"Lang", "Term", "Definition", []Node{
	{"Lang", "descendant", "", []Node{}},
	{"Lang", "progeny", "", []Node{}},
}}

func TestNodeHtml(t *testing.T) {
	// Space after span closing tag is annoying
	expected := "<li><span complete><h1>Lang</h1><h2>Term</h2><p>Definition</p></span> <ul><li><span complete><h1>Lang</h1><h2>descendant</h2></span> </li><li><span complete><h1>Lang</h1><h2>progeny</h2></span> </li></ul></li>"

	buf := bytes.NewBuffer([]byte{})
	testNode.html().Render(context.Background(), buf)
	out := buf.String()

	if out != expected {
		t.Errorf("input: '%v', expected output '%s' but got '%s'",
			testTree, expected, out)
	}
}

func TestNodeObscure(t *testing.T) {
	obscurer := '_'
	tests := []struct {
		in       []string
		expected Node
	}{
		{[]string{},
			Node{"Lang", "T___", "", []Node{
				{"Lang", "d_________", "", []Node{}},
				{"Lang", "p______", "", []Node{}},
			}},
		},
		{[]string{"term", "descendant", "progeny"}, testNode},
		{[]string{"term", "descendant", "pro"},
			Node{"Lang", "Term", "Definition", []Node{
				{"Lang", "descendant", "", []Node{}},
				{"Lang", "pro____", "", []Node{}},
			}},
		},
		{[]string{"term", "descendant"},
			Node{"Lang", "Term", "Definition", []Node{
				{"Lang", "descendant", "", []Node{}},
				{"Lang", "p______", "", []Node{}},
			}},
		},
	}

	for _, test := range tests {
		out := testNode.obscure(test.in, obscurer)
		if reflect.DeepEqual(out, test.expected) {
			t.Errorf("input: '%v', expected output '%v' but got '%v'", test.in, test.expected, out)
		}
	}
}

func TestNodeComplete(t *testing.T) {
	obscurer := '_'
	tests := []struct {
		in       Node
		expected bool
	}{
		{testNode, true},
		{testNode.obscure([]string{}, obscurer), false},
		{testNode.obscure([]string{"term", "descendant", "progeny"}, obscurer), true},
		{testNode.obscure([]string{"term", "descendant"}, obscurer), false},
	}

	for _, test := range tests {
		out := test.in.isComplete(obscurer)
		if out != test.expected {
			t.Errorf("input: '%v', expected output '%t' but got '%t'", test.in, test.expected, out)
		}
	}
}
