package app

import "github.com/aditya43/golang-bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
