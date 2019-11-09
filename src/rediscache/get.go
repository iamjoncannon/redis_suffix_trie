package rediscache

import (
	"github.com/gomodule/redigo/redis"
)

// GetCache queries the connection pool
func GetCache(c redis.Conn, key string) (s string, e error) {

	s, err := redis.String(c.Do("LRANGE", key, "0 -1"))

	return s, err
}
