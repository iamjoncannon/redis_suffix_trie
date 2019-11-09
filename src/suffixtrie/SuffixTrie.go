package suffixtrie

import (
	"encoding/json"
	"io/ioutil"
	"rediscache"
)

/*

we want the trie to be a map with strings as a key
and a value that could either be *Trie, or a
payload, map[string]string

in other languages, this would be a union or sum type,
in Go we can do this a few ways

https://making.pusher.com/alternatives-to-sum-types-in-go/

*/

// TrieNode are the possible ends of the trie hash
type TrieNode interface {
	// any type that implements isTrieNode method
	// is a TrieNode
	isTrieNode() string
}

// Trie is the trie object
type Trie map[string]TrieNode

// EndSymbol denotes the payload on the TrieNode
func (t Trie) EndSymbol() string {

	return "#"
}

func (t Trie) isTrieNode() string {

	return "Its a Trie"
}

// NewTrie creates a new Trie
func NewTrie() *Trie {

	return &Trie{}
}

// Payload are all the items that contain a substring that
// ends at that specific node-
type Payload map[string][]string

func (p Payload) isTrieNode() string {

	return "Its a Payload"
}

// init function
func newPayload() *Payload {

	return &Payload{}
}

// PrintTrieToFile generates a JSON and
// writes to file
func (t *Trie) PrintTrieToFile(name string) {

	data, err := json.Marshal(t)

	if err != nil {

		panic(err)
	}

	err = ioutil.WriteFile(name+".json", data, 0644)

	if err != nil {

		panic(err)
	}
}

// InsertIntoTrie takes each individual string
// and inserts in into the entire Trie structure
func (t *Trie) InsertIntoTrie(item string, payload string) {

	pool := rediscache.NewPool()
	conn := pool.Get()
	defer conn.Close()
	err := rediscache.Ping(conn)

	if err != nil {

		panic(err)
	}

	endSymbol := (*t).EndSymbol()

	for i := range item {

		var node *Trie
		node = t

		currentPath := ""

		for j := i; j < len(item); j++ {

			letter := string(item[j])

			currentPath += letter

			// this node doesn't exist yet

			if (*node)[letter] == nil {

				nextTrie := NewTrie()

				(*node)[letter] = nextTrie
			}

			node = (*node)[letter].(*Trie)

			// skip the first level

			if node != t {

				payloadNode, hasPayloadNode := (*node)[endSymbol].(*Payload)

				// there's already a payload
				// on this node
				// example-- "aba" "bab"
				if hasPayloadNode == true {

					_, hasThisPayloadAlready := (*payloadNode)[payload]

					// the payload node is a hash to prevent duplicate entries--
					if hasThisPayloadAlready == false {

						(*payloadNode)[payload] = append((*payloadNode)[payload], payload)

						rediscache.SetCache(conn, currentPath, payload)
					}

				} else {

					(*node)[endSymbol] = newPayload()

					payloadNode := (*node)[endSymbol].(*Payload)

					(*payloadNode)[payload] = payload
					rediscache.SetCache(conn, currentPath, payload)
				}

			} else {

				// this is a node at the root of the trie
				// we want to store the payload, but not
				// at the top level of the trie, but on
				// this node

				// so we advance to next level before
				// storing the payload

				(*node)[endSymbol] = newPayload()

				payloadNode := (*node)[endSymbol].(*Payload)

				(*payloadNode)[payload] = payload

				rediscache.SetCache(conn, currentPath, payload)
			}
		}
	}
}
