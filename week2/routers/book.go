package routers

import (
	"github.com/anaghamalkar08/2-GL1-CipherSchools/database"
	"github.com/anaghamalkar08/2-GL1-CipherSchools/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default() //default engine
	api := handler.Handler{
		DB: database.GetDB(), //set handler DB
	}
	router.GET("/books", api.GetBooks) //set func
	router.POST("/books", api.SaveBook)
	return router
}
