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

	thisTrie.InsertIntoTrieRedis("The rain in spain falls mainly in the plain", "payloadd", conn)

}
