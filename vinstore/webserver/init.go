package webserver

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

// Start registers the request handlers for the webserver and starts serving
// requests.
func Start(addr string) error {
	r := httprouter.New()

	r.GET("/api/matches", Matches)
	r.GET("/api/games", Games)
	r.GET("/api/new_games", NewGames)

	return http.ListenAndServe(addr, r)
}

// DB is the gorm database that is going to be used for requesting information
// from the database.
var DB *gorm.DB
