package app

import (
	"github.com/aditya43/golang-bookstore_users-api/logger"
	"github.com/aditya43/golang-bookstore_users-api/utils/env"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	gin.SetMode(env.Get("GIN_MODE"))
	mapUrls()

	logger.Info("Application starting...")
	if err := router.Run(":8080"); err != nil {
		logger.Error(err)
		panic(err)
	}
}
