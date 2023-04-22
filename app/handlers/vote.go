package handlers

import (
	"Voting-API/db"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Vote struct {
	Name string `json:"name"`
}

//need fix and update
func CreateVote(conn redis.Conn, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vote Vote
	err := decoder.Decode(&vote)
	if err != nil {
		fmt.Fprintf(w, "nooooy")
		return
	}
	_, err = db.AddVote(conn, vote.Name)

	fmt.Fprintf(w, "vote added!")
}

func GetVote(conn redis.Conn, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nameVote := vars["name"]
	result, _ := redis.StringMap(conn.Do("HGETALL", nameVote))
	jsonResult, _ := json.Marshal(result)
	fmt.Fprintf(w, string(jsonResult))
}

func IncrementVote(conn redis.Conn, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nameVote := vars["name"]
	value, _ := redis.Int(conn.Do("HINCRBY", nameVote, "count", 1))
	fmt.Fprintf(w, strconv.Itoa(value))
}
