package main

import (
	"fmt"
	"rediscache"
	"suffixtrie"
)

func main() {

	pool := rediscache.NewPool()
	conn := pool.Get()
	defer conn.Close()

	thisTrie := suffixtrie.NewTrie()

	thisTrie.InsertIntoTrieRedis("The rain in spain falls mainly in the plain", "payloadd", conn)
	thisTrie.InsertIntoTrieRedis("The rain in spain falls mainly in the plain", "another payloaaddd", conn)

	output := thisTrie.ContainsRedis("a", conn)

	fmt.Println(output)

}
