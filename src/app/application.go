package app

import (
	"github.com/LordRadamanthys/bookstore_oauth-api/src/domain/access_token"
	"github.com/LordRadamanthys/bookstore_oauth-api/src/domain/access_token/repository/db"
	"github.com/LordRadamanthys/bookstore_oauth-api/src/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8082")
}
