package main

import (
	"github.com/NDoolan360/etyml-app/web/templates"
	"github.com/a-h/templ"
)

type Node struct {
	Lang       string `json:"language"`
	Term       string `json:"term"`
	Definition string `json:"definition"`
	Children   []Node `json:"children"`
}

func (node Node) html() templ.Component {
	var children []templ.Component
	for _, child := range node.Children {
		newChild := child.html()
		children = append(children, newChild)
	}

	goal := node.Lang == "English"
	complete := node.isComplete('_')

	return templates.Node(
		node.Lang,
		node.Term,
		node.Definition,
		children,
		goal,
		complete,
	)
}

func (node Node) obscure(guesses []string, obscurer rune) Node {
	var newChildren []Node
	shouldObscure := true
	for _, child := range node.Children {
		newChild := child.obscure(guesses, obscurer)
		if isCompletelyUnobscured(newChild.Term, obscurer) {
			// on full reveal of child, reveal parents
			shouldObscure = false
		}
		newChildren = append(newChildren, newChild)
	}

	if shouldObscure {
		node.Term = obscureTerm(node.Term, guesses)
	}

	return Node{
		node.Lang,
		node.Term,
		"", // remove description
		newChildren,
	}
}

func obscureTerm(term string, guesses []string) string {
	hideIndex := largestGuess(term, guesses)
	obscuredWithGuesses := obscureStringAfterNth(term, hideIndex, '_')
	if isCompletelyObscured(obscuredWithGuesses, '_') {
		// sneak peek
		return obscureStringAfterNth(term, 1, '_')
	} else {
		return obscuredWithGuesses
	}
}

func (node Node) isComplete(obscurer rune) bool {
	childrenObscured := false
	for _, child := range node.Children {
		if !child.isComplete(obscurer) {
			childrenObscured = true
			break
		}
	}
	return isCompletelyUnobscured(node.Term, obscurer) && !childrenObscured
}
