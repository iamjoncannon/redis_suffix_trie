package main

import (
	"fmt"
	"suffixtrie"
)

func main() {

	thisTrie := suffixtrie.NewTrie()

	thisTrie.InsertIntoTrie("the rain in spain falls mainly in the plains", "123")
	thisTrie.InsertIntoTrie("the rain in spain falls mainly in the plains", "456")

	fmt.Println((*thisTrie)["a"])
}
