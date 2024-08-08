package main

import (
	"bytes"
	"context"
	"reflect"
	"testing"
)

var testTree = Tree{
	Node{"jnvoiuhef", "Lang", "Term", "Definition",
		[]Node{
			{"noiusdnml", "Lang", "*child", "", []Node{}},
			{"lkjdafohc", "Lang", "children", "", []Node{}},
		},
	},
}

func TestTreeHtml(t *testing.T) {
	// Space after span closing tag is annoying
	expected := `<ul><li id="jnvoiuhef"><span complete><h1>Lang</h1><h2>Term</h2><p>Definition</p></span> <ul><li id="noiusdnml"><span complete><h1>Lang</h1><h2>*child</h2></span> </li><li id="lkjdafohc"><span complete><h1>Lang</h1><h2>children</h2></span> </li></ul></li></ul>`

	buf := bytes.NewBuffer([]byte{})
	testTree.html([]string{}).Render(context.Background(), buf)
	out := buf.String()

	if out != expected {
		t.Errorf("input: '%v', expected output '%s' but got '%s'",
			testTree, expected, out)
	}
}

func TestTreeObscure(t *testing.T) {
	obscurer := '_'
	tests := []struct {
		in       []string
		expected Tree
	}{
		{[]string{}, Tree{
			Node{"jnvoiuhef", "Lang", "Term", "Definition",
				[]Node{
					{"noiusdnml", "Lang", "c____", "", []Node{}},
					{"lkjdafohc", "Lang", "c_______", "", []Node{}},
				},
			},
		}},
		{[]string{"term", "child", "children"}, testTree},
		{[]string{"term", "child"}, Tree{
			Node{"jnvoiuhef", "Lang", "Term", "Definition",
				[]Node{
					{"noiusdnml", "Lang", "child", "", []Node{}},
					{"lkjdafohc", "Lang", "child___", "", []Node{}},
				},
			},
		}},
	}

	for _, test := range tests {
		out := testTree.obscure(test.in, []string{}, obscurer)
		if reflect.DeepEqual(out, test.expected) {
			t.Errorf("input: '%v', expected output '%v' but got '%v'", test.in, test.expected, out)
		}
	}
}

func TestTreeComplete(t *testing.T) {
	obscurer := '_'
	tests := []struct {
		in       Tree
		expected bool
	}{
		{testTree, true},
		{testTree.obscure([]string{}, []string{}, obscurer), false},
		{testTree.obscure([]string{"term", "child", "children"}, []string{}, obscurer), true},
		{testTree.obscure([]string{"term", "child"}, []string{}, obscurer), false},
	}

	for _, test := range tests {
		out := test.in.isComplete(obscurer)
		if out != test.expected {
			t.Errorf("input: '%v', expected output '%t' but got '%t'", test.in, test.expected, out)
		}
	}
}
