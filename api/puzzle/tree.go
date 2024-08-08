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

var etymologyTrees = map[string]Tree{
	"M91vw3KIWzmaNuSNk5453Q": {
		Node{
			"aefnilwnf", "Proto-Indo-European", "*men-", "think", []Node{
				{"awdboengz", "Late Latin", "mentālis", "", []Node{
					{"wnhsuifbg", "Middle French", "mental", "", []Node{
						{"snwifnjkn", "English", "mental", "Of or relating to intellectual as contrasted with emotional activity.", []Node{}},
					}},
				}},
				{"dbfjsnkfh", "Proto-Indo-Iranian", "*mántram", "", []Node{
					{"dnkauwnwm", "Proto-Indo-Aryan", "*mántram", "", []Node{
						{"wnjfusnal", "Sanskrit", "mantra (मन्त्र)", "", []Node{
							{"mnbhyujsw", "English", "mantra", "The hymn portions of the Vedas; any passage of these used as a prayer.", []Node{}},
							{"plsmkhdbi", "Malay", "menteri", "", []Node{
								{"snikwmali", "Portuguese", "mandarim", "", []Node{
									{"mfndhuwje", "English", "mandarin", "A pedantic or elitist bureaucrat.", []Node{}},
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
			"wqmksubdn", "Proto-Indo-European", "*bʰer-", "to carry, bear", []Node{
				{"fodnsbuej", "Proto-Germanic", "*barô", "", []Node{
					{"snibekcjw", "Frankish", "*barō", "", []Node{
						{"plwmwnehf", "Medieval Latin", "barō", "", []Node{
							{"sjiwsncju", "Old French", "baron", "", []Node{
								{"snhwihdin", "Middle English", "baroun", "", []Node{
									{"smkallwwe", "English", "baron", "A male member of the lowest rank of English nobility.", []Node{}},
								}},
							}},
						}},
					}},
				}},
				{"plwnjjisb", "Latin", "fūr", "", []Node{
					{"sohdbjwku", "Vulgar Latin", "*furittum", "", []Node{
						{"wkkduisbb", "Old French", "furet", "", []Node{
							{"aosjfbeek", "Middle English", "furet", "", []Node{
								{"sjiebfjss", "English", "ferret", "An often domesticated mammal rather like a weasel, descended from the polecat and often trained to hunt burrowing animals.", []Node{}},
							}},
						}},
					}},
					{"flrirnjii", "Latin", "sufferō", "", []Node{
						{"dloeebbgk", "Anglo-Norman", "suffrir", "", []Node{
							{"wkoplenwk", "Middle English", "suffren", "", []Node{
								{"njoplednh", "English", "suffer", "To undergo hardship.", []Node{}},
							}},
						}},
					}},
				}},
			},
		},
	},
	"EMuXo_ueWy2M3HuxvCp70A": {
		Node{
			"snkolwkdn", "Proto-Indo-European", "*h₂er-", "to fit, to fix, to put together, to slot", []Node{
				{"nwodlibee", "Latin", "arma", "", []Node{
					{"eebkjodnp", "Middle French", "alarme", "", []Node{
						{"kdnjwolkm", "Middle English", "alarom", "", []Node{
							{"wjddplenk", "Middle English", "alarme", "", []Node{
								{"dlinsmkwl", "English", "alarm", "A summons to arms, as on the approach of an enemy.", []Node{}},
							}},
						}},
					}},
				}},
				{"kwmkfobtj", "Latin", "reor", "", []Node{
					{"hvukjsbjj", "Latin", "ratiō", "", []Node{
						{"fbwjuebni", "French", "ration", "", []Node{
							{"dknvhufir", "English", "ration", "To portion out (especially during a shortage of supply); to limit access to.", []Node{}},
						}},
					}},
				}},
				{"kfngbouen", "Latin", "ornāre", "", []Node{
					{"dnigktnei", "Latin", "ornamentum", "", []Node{
						{"ifkrnwpql", "Old French", "ornement", "", []Node{
							{"dkwofnghr", "Middle English", "ornament", "", []Node{
								{"dnjeogntk", "English", "ornament", "An element of decoration; that which embellishes or adorns.", []Node{}},
							}},
						}},
					}},
				}},
			},
		},
	},
}
