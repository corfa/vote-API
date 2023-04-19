package redisQ

import "github.com/gomodule/redigo/redis"

func AddVote(conn redis.Conn, name string) interface{} {

	res, _ := conn.Do("HSET", name, "count", 0)

	return res

}
