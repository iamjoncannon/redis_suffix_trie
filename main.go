package main

import (
	"rediscache"
	"suffixtrie"
)

func main() {

	pool := rediscache.NewPool()
	conn := pool.Get()
	defer conn.Close()

	thisTrie := suffixtrie.NewTrie()

	thisTrie.InsertIntoTrieRedis("The rain in spain falls mainly in the plain", "payload", conn)
	thisTrie.InsertIntoTrieRedis("spain", "another payload", conn)
	thisTrie.InsertIntoTrieRedis("plain", "a third payload", conn)

	queryOne := thisTrie.ContainsRedis("rain", conn)  // ["payload"]
	queryOne := thisTrie.ContainsRedis("spain", conn) // ["payload", "another payload"]
	queryOne := thisTrie.ContainsRedis("plain", conn) // ["payload", "a third payload"]
	outputTwo := thisTrie.ContainsRedis("zoo", conn)  // <NOT FOUND>

}
