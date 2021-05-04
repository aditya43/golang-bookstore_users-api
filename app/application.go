package app

import (
	"github.com/aditya43/golang-bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Log.Info("Application started. Listening on port 8080")
	_ = router.Run(":8080")
}
