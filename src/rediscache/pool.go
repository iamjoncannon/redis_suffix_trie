package rediscache

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

/*

code to instantiate pool:

pool := rediscache.NewPool()
conn := pool.Get()
defer conn.Close()

conn is the connection object

*/

func NewPool() *redis.Pool {

	return &redis.Pool{ // Maximum number of idle connections in the pool.
		MaxIdle: 10000,
		// max number of connections
		MaxActive:   12000,
		IdleTimeout: 240 * time.Second,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())

			}

			return c, err
		},
	}
}
