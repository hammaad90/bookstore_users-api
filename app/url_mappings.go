package app

import (
	"github.com/hammaad90/bookstore_users-api/controllers/ping"
	"github.com/hammaad90/bookstore_users-api/controllers/users"
)

func mapUrls() {

	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	// router.GET(relativePath: "/users/search", controllers.SearchUser)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)

	// patch is used for partial update
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
