package app

import (
	"github.com/Resoluter/bookstore-users.git/controllers/ping"
	"github.com/Resoluter/bookstore-users.git/controllers/users"
)

func mapUrl() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
