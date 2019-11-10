# Redis Suffix Trie

"A suffix trie contains all the suffixes of the given text as their keys and positions in the text as their values. Suffix tries allow particularly fast implementations of string operations."

https://en.wikipedia.org/wiki/Suffix_tree

This package implements the trie data structure in Redis by storing each suffix as the key to a list, containing all values/payloads inserted for that suffix. Unlike some tries, its a "shallow" hash- only one dimension, since Redis doesn't natively support nested data structures. 

Potential uses cases-- indexing a series of files for quick searching, indexing occurences of terms in a text. The API could also be modified easily to support counting occurances. 

# API

## Redis methods

```golang

func main() {

	pool := rediscache.NewPool()
	conn := pool.Get()
	defer conn.Close()

    thisTrie := suffixtrie.NewTrie()
    
    // Inserts the payload for the strings and substrings of the term
    thisTrie.InsertIntoTrieRedis("supercalifragilisticexpialidocious", "you'll always sound precocious", conn)
    thisTrie.InsertIntoTrieRedis("super", "But then one day I learned a word.", conn)
    thisTrie.InsertIntoTrieRedis("califrag", "The biggest word you ever heard. And this is how it goes", conn)

    // Query the cache for any string inserted into the trie, and receive the payload
    // stored at each inserted string  
    thisTrie.ContainsRedis("docious", conn)  // ["you'll always sound precocious"]
    thisTrie.ContainsRedis("super", conn) // ["you'll always sound precocious", "But then one day I learned a word."]
    thisTrie.ContainsRedis("califrag", conn) // ["you'll always sound precocious", "a third payload"]
    thisTrie.ContainsRedis("Merry Poppins", conn)  // <NOT FOUND>
}
```

## Without Redis

The library also supports using the trie structure without storing in Redis cache 

```golang

thisTrie := suffixtrie.NewTrie()

thisTrie.InsertIntoTrie("supercalifragilisticexpialidocious", "even though the sound of it is really quite atrocious", conn)

queryOne := thisTrie.Contains("supercalifragilisticexpialidocious")  // ["even though the sound of it is really quite atrocious"]

```

### PrintTrieToFile

Outputs JSON

```golang

thisTrie.PrintTrieToFile("MerryPoppinsLyrics") // MerryPoppinsLyrics.json

```
