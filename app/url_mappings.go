package app

import (
	"moviestore_users-api/controllers/ping"
	"moviestore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.POST("/users/search", users.Search)
}
