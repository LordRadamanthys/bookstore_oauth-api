package app

import (
	"github.com/LordRadamanthys/bookstore_oauth-api/controllers/ping"
	"github.com/LordRadamanthys/bookstore_oauth-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
}
