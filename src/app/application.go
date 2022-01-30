package app

import (
	"github.com/LordRadamanthys/bookstore_oauth-api/src/http"
	"github.com/LordRadamanthys/bookstore_oauth-api/src/repository/db"
	"github.com/LordRadamanthys/bookstore_oauth-api/src/repository/rest"
	"github.com/LordRadamanthys/bookstore_oauth-api/src/services/access_token_service"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	atHandler := http.NewHandler(access_token_service.NewService(rest.NewRestRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
