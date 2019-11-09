package main

import (
	"suffixtrie"
)

func main() {

	thisTrie := suffixtrie.NewTrie()

	// thisTrie.HydrateRedisCache()
	thisTrie.InsertIntoTrie("Aphex Twin- XTal 290134- genre- ambient", "Aphex")
	thisTrie.InsertIntoTrie("The Taurii have weapons of the ancients", "SG1")

	// thisTrie.PrintTrie("AphexTrie")

	// thisTrie.LoadTrie("AphexTrie.json")
}
