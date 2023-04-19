package handlers

import (
	"Voting-API/redisQ"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"net/http"
)

type Vote struct {
	Name string `json:"name"`
}

func CreateVote(conn redis.Conn, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vote Vote
	err := decoder.Decode(&vote)
	if err != nil {
		// обработка ошибки
	}
	redisQ.AddVote(conn, vote.Name)
	fmt.Fprintf(w, "vote added!")
}
