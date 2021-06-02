package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hammaad90/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start deamon...")
	router.Run(":8080")
}
