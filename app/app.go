package app

import (
	"Voting-API/app/handlers"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// App has router and db instances
type App struct {
	Router    *mux.Router
	RedisConn redis.Conn
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//ДЛЯ РЕДИСА
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	//var redisDbNumber int = os.Getenv("REDIS_DB_NUMBER")
	a.RedisConn, _ = redis.Dial("tcp", redisHost+":"+redisPort)

	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Post("/vote", a.handleRequest(handlers.CreateVote))
	a.Post("/vote/{name}", a.handleRequest(handlers.IncrementVote))
	a.Get("/vote/{name}", a.handleRequest(handlers.GetVote))

}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(connRedis redis.Conn, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.RedisConn, w, r)
	}
}
