package main

import (
	"github.com/NDoolan360/etyml-app/web/templates"
	"github.com/a-h/templ"
)

type Node struct {
	Id         string `json:"id"`
	Lang       string `json:"language"`
	Term       string `json:"term"`
	Definition string `json:"definition"`
	Children   []Node `json:"children"`
}

func (node Node) html(hints []string) templ.Component {
	var children []templ.Component
	for _, child := range node.Children {
		newChild := child.html(hints)
		children = append(children, newChild)
	}

	goal := node.Lang == "English"
	complete := node.isComplete('_')

	return templates.Node(
		node.Id,
		node.Lang,
		node.Term,
		node.Definition,
		children,
		goal,
		complete,
		node.alreadyHinted(hints),
	)
}

func (node Node) obscure(guesses []string, hints []string, obscurer rune) Node {
	var newChildren []Node
	shouldObscure := true
	for _, child := range node.Children {
		newChild := child.obscure(guesses, hints, obscurer)
		if isCompletelyUnobscured(newChild.Term, obscurer) {
			// on full reveal of child, reveal parents
			shouldObscure = false
		}
		newChildren = append(newChildren, newChild)
	}

	node.Children = newChildren

	if shouldObscure {
		node.Term = obscureTerm(node.Term, guesses)
		if !(node.alreadyHinted(hints) || node.isComplete(obscurer)) {
			node.Definition = ""
		}
	}

	return node
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

func (node Node) alreadyHinted(hints []string) bool {
	for _, hint := range hints {
		if node.Id == hint {
			return true
		}
	}
	return false
}
