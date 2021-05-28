package app

import (
	"github.com/hammaad90/bookstore_users-api/controllers/ping"
	"github.com/hammaad90/bookstore_users-api/controllers/users"
)

func mapUrls() {

	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	// router.GET(relativePath: "/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}
