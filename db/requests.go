package db

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)

func AddVote(conn redis.Conn, name string) (interface{}, error) {

	res, err := conn.Do("HSET", name, "count", 0)
	if err != nil {
		return nil, err
	}
	return res, err

}

func GetVote(conn redis.Conn, name string) ([]byte, error) {
	result, err := redis.StringMap(conn.Do("HGETALL", name))
	jsonResult, _ := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return jsonResult, err
}
