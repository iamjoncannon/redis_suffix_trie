package suffixtrie

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertIntoTrie(t *testing.T) {

	firstTrie := NewTrie()

	firstTrie.InsertIntoTrie("a", "123")

	if (*firstTrie).isTrieNode() != "Its a Trie" {

		t.Errorf("firstTrie is not a Trie")
	}

	if reflect.TypeOf(firstTrie).String() != "*suffixtrie.Trie" {

		t.Errorf("firstTrie is not a Trie")
	}

	aNode := (*firstTrie)["a"]

	if aNode.isTrieNode() != "Its a Trie" {

		t.Errorf("a is not a Trie")
	}

	if _, ok := aNode.(*Trie); ok != true {

		t.Errorf("a is not a Trie")
	}

	if payload, ok := (*aNode.(*Trie))["#"]; ok != true {

		fmt.Println(firstTrie, aNode, payload, ok)
		t.Errorf("a trie has no payload")
	}

}
