package suffixtrie

import (
	"encoding/json"
	"io/ioutil"
	"rediscache"

	"github.com/gomodule/redigo/redis"
)

// this is the second version-- same api, but uses a
// one dimensional hash

// Trie is a flat hash map
type Trie map[string][]string

// NewTrie is the constructor/__init__ method
func NewTrie() *Trie {

	return &Trie{}
}

// PrintTrieToFile prints a trie to JSON
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

// InsertIntoTrie inserts into a trie in the running application memory
func (t *Trie) InsertIntoTrie(item string, payload string) {

	for i := 0; i < len(item); i++ {

		for j := i + 1; j < len(item)+1; j++ {

			chunk := string(item[i:j])

			thisEntry, ok := (*t)[chunk]

			if ok != true {

				(*t)[chunk] = []string{payload}

			} else {

				var isAlreadyInPayload bool

				isAlreadyInPayload = false

				for _, currentPayload := range thisEntry {

					if currentPayload == payload {

						isAlreadyInPayload = true
					}
				}

				if isAlreadyInPayload == false {

					(*t)[chunk] = append(thisEntry, payload)
				}
			}
		}
	}

}

// InsertIntoTrieRedis inserts items into Redis trie rather than
// application memory
func (t *Trie) InsertIntoTrieRedis(item string, payload string, conn redis.Conn) {

	for i := 0; i < len(item); i++ {

		for j := i + 1; j < len(item)+1; j++ {

			chunk := string(item[i:j])

			// query Redis for this key
			thisEntry, ok := rediscache.GetCache(conn, chunk)

			/*
				if ok != true {

					(*t)[chunk] = []string{payload}

				} else {

					var isAlreadyInPayload bool

					isAlreadyInPayload = false

					for _, currentPayload := range thisEntry {

						if currentPayload == payload {

							isAlreadyInPayload = true
						}
					}

					if isAlreadyInPayload == false {

						(*t)[chunk] = append(thisEntry, payload)
					}
				}

			*/
		}
	}
}

// Contains returns currently stored value, or <NOT FOUND>
// as []string
func (t *Trie) Contains(item string) []string {

	thisAttempt, ok := (*t)[item]

	if ok != true {

		return []string{"<NOT FOUND>"}
	}

	return thisAttempt
}
