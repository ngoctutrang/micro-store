package app

import (
	"github.com/ngoctutrang/micro-store/controllers/ping"
	"github.com/ngoctutrang/micro-store/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)

}
