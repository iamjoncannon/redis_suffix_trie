package rediscache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func Ping(c redis.Conn) error {
	s, err := redis.String(c.Do("PING"))

	if err != nil {
		return err
	}

	fmt.Printf("PING Response = redis says %s\n'", s)

	return nil
}
