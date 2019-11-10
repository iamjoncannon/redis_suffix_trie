package rediscache

import (
	"github.com/gomodule/redigo/redis"
)

// GetCache queries the connection pool
func GetCache(c redis.Conn, key string) (s interface{}, e error) {

	items, err := redis.Strings(c.Do("LRANGE", key, "0", "-1"))
	return items, err
}
