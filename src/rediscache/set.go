package rediscache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// RPUSH a 'thang dose'

// SetCache sets a list item with rpush
func SetCache(c redis.Conn, key string, value string) error {

	_, err := c.Do("RPUSH", key, value)

	if err != nil {
		fmt.Println("rpush error", err)
	}

	return err
}
