package main

import (
	"github.com/NDoolan360/etyml-app/web/templates"
	"github.com/a-h/templ"
)

type Tree struct {
	Etymologies Node `json:"etymology"`
}

func (tree Tree) html() templ.Component {
	return templates.Tree(tree.Etymologies.html())
}

func (tree Tree) obscure(guesses []string) Tree {
	return Tree{
		Node{
			tree.Etymologies.Lang,
			tree.Etymologies.Term,
			tree.Etymologies.Definition,
			tree.Etymologies.obscure(guesses).Children,
		},
	}
}

func (tree Tree) isComplete() bool {
	return tree.Etymologies.isComplete('_')
}

var etymologyTrees = map[string]Tree{
	"M91vw3KIWzmaNuSNk5453Q": {
		Node{
			"Proto-Indo-European", "*men-", "think", []Node{
				{"Late Latin", "mentālis", "", []Node{
					{"Middle French", "mental", "", []Node{
						{"English", "mental", "", []Node{}},
					}},
				}},
				{"Proto-Indo-Iranian", "*mántram", "", []Node{
					{"Proto-Indo-Aryan", "*mántram", "", []Node{
						{"Sanskrit", "mantra (मन्त्र)", "", []Node{
							{"English", "mantra", "", []Node{}},
							{"Malay", "menteri", "", []Node{
								{"Portuguese", "mandarim", "", []Node{
									{"English", "mandarin", "", []Node{}},
								}},
							}},
						}},
					}},
				}},
			},
		},
	},
	"pJgLI2smVVyrmh2dZdM0cg": {
		Node{
			"Proto-Indo-European", "*bʰer-", "to carry, bear", []Node{
				{"Proto-Germanic", "*barô", "", []Node{
					{"Frankish", "*barō", "", []Node{
						{"Medieval Latin", "barō", "", []Node{
							{"Old French", "baron", "", []Node{
								{"Middle English", "baroun", "", []Node{
									{"English", "baron", "", []Node{}},
								}},
							}},
						}},
					}},
				}},
				{"Latin", "fūr", "", []Node{
					{"Vulgar Latin", "*furittum", "", []Node{
						{"Old French", "furet", "", []Node{
							{"Middle English", "furet", "", []Node{
								{"English", "ferret", "", []Node{}},
							}},
						}},
					}},
					{"Latin", "sufferō", "", []Node{
						{"Anglo-Norman", "suffrir", "", []Node{
							{"Middle English", "suffren", "", []Node{
								{"English", "suffer", "", []Node{}},
							}},
						}},
					}},
				}},
			},
		},
	},
	"EMuXo_ueWy2M3HuxvCp70A": {
		Node{
			"Proto-Indo-European", "*h₂er-", "to fit, to fix, to put together, to slot", []Node{
				{"Latin", "arma", "", []Node{
					{"Middle French", "alarme", "", []Node{
						{"Middle English", "alarom", "", []Node{
							{"Middle English", "alarme", "", []Node{
								{"English", "alarm", "", []Node{}},
							}},
						}},
					}},
				}},
				{"Latin", "reor", "", []Node{
					{"Latin", "ratiō", "", []Node{
						{"French", "ration", "", []Node{
							{"English", "ration", "", []Node{}},
						}},
					}},
				}},
				{"Latin", "ornāre", "", []Node{
					{"Latin", "ornamentum", "", []Node{
						{"Old French", "ornement", "", []Node{
							{"Middle English", "ornament", "", []Node{
								{"English", "ornament", "", []Node{}},
							}},
						}},
					}},
				}},
			},
		},
	},
}
