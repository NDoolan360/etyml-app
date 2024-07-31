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

func (node Node) obscure(guesses []string) Node {
	obscuredTerm := obscureTerm(node.Term, guesses)

	var newChildren []Node
	shouldObscure := true
	for _, child := range node.Children {
		newChild := child.obscure(guesses)
		if isCompletelyUnobscured(newChild.Term, '_') {
			// on full reveal of child, reveal parents
			shouldObscure = false
		}
		newChildren = append(newChildren, newChild)
	}

	if shouldObscure {
		return Node{
			node.Lang,
			obscuredTerm,
			"",
			newChildren,
		}
	} else {
		return Node{
			node.Lang,
			node.Term,
			"",
			newChildren,
		}
	}
}

func obscureTerm(term string, guesses []string) string {
	hideIndex := largestGuess(term, guesses)
	obscuredWithGuesses := obscureStringAfterNth(term, hideIndex, '_')
	if isCompletelyObscured(obscuredWithGuesses, '_') {
		return sneakPeek(term, '_')
	} else {
		return obscuredWithGuesses
	}
}

func (node Node) isComplete(obscurer rune) bool {
	chilrenObscured := false
	for _, child := range node.Children {
		if !child.isComplete(obscurer) {
			chilrenObscured = true
			break
		}
	}
	return isCompletelyUnobscured(node.Term, obscurer) && !chilrenObscured
}
