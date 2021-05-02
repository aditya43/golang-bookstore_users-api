package app

import (
	"github.com/aditya43/golang-bookstore_users-api/controllers/ping"
	"github.com/aditya43/golang-bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Pong)

	router.GET("/users/search", users.Search)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.POST("/users", users.Create)
}
