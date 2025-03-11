package main

import (
	"github.com/NDoolan360/etyml-app/web/templates"
	"github.com/a-h/templ"
)

type Tree struct {
	Etymologies Node `json:"etymology"`
}

func (tree Tree) html(hints []string) templ.Component {
	return templates.Tree(tree.Etymologies.html(hints))
}

func (tree Tree) obscure(guesses []string, hints []string, obscurer rune) Tree {
	return Tree{
		Node{
			tree.Etymologies.Id,
			tree.Etymologies.Lang,
			tree.Etymologies.Term,
			tree.Etymologies.Definition,
			tree.Etymologies.obscure(guesses, hints, obscurer).Children,
		},
	}
}

func (tree Tree) isComplete(obscurer rune) bool {
	return tree.Etymologies.isComplete(obscurer)
}
