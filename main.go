package main

import (
	"api-translate/controller"
	"api-translate/docs"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
	"github.com/subosito/gotenv"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	if err := gotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

// @title           API TRANSLATE
// @version         1.0
// @description     Service For API TRANSLATE
func main() {
	baseApi := "/api/v1"
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = baseApi

	apiV1 := router.Group(baseApi)
	controller.NewTranslateTextController(apiV1)
	apiV1.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Imusegipo"}) })
	log.Println(aurora.Green(fmt.Sprintf("Server running on http://localhost%s/swagger/index.html", port)))
	router.Run(port)
}
