package rediscache

import (
	"github.com/gomodule/redigo/redis"
)

func GetCache(c redis.Conn, key string) (s string, e error) {

	s, err := redis.String(c.Do("hget", "trie", key))

	return s, err
}
