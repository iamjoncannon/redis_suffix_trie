package rediscache

import (
	"github.com/gomodule/redigo/redis"
)

func SetCache(c redis.Conn, key string, value string) {

	_, err := c.Do("hset", "trie", key, value)

	if err != nil {
		panic(err)
	}
}
