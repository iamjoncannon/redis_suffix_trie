package rediscache

import (
	"github.com/gomodule/redigo/redis"
)

func GetCache(c redis.Conn, key string, value string) string {

	s, err := redis.String(c.Do("hget", "trie", key, value))

	return s, err

	/*
		return func(w http.ResponseWriter, r *http.Request) {

			allParams := mux.Vars(r)

			param := allParams["key"]

			s, err := redis.String(c.Do("GET", param))

			returnJSON := make(map[string]string)

			if err == redis.ErrNil {

				fmt.Printf("%s does not exist\n", param)

			} else if err != nil {

				fmt.Println(err)
			} else {

				returnJSON[param] = s
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(returnJSON)
			}
		}
	*/
}
