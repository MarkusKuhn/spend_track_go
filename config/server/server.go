package server

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"spend_track_go/api/transactions"
	"spend_track_go/helpers/utilities"
)

func Start() {
	var port string = utilities.GetEnvVariable("SERVER_PORT")
	var portString string = fmt.Sprintf(":%s", port)

	router := initializeNewRouter()

	log.Print(fmt.Sprintf("Running on port: %s", portString))
	router.Run(portString)
}

func initializeNewRouter() *gin.Engine {
	router := gin.Default()

	configureMiddleware(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "index"})
	})

	router.GET("/transactions", func(c *gin.Context) {
		c.JSON(200, gin.H{"transactions": transactions.GetTransactions()})
	})

	return router
}

func configureMiddleware(router *gin.Engine) {
	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"http://localhost"},
				AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
				AllowHeaders:     []string{"Origin"},
				ExposeHeaders:    []string{"Content-Length"},
				AllowCredentials: true,
				AllowOriginFunc: func(origin string) bool {
					return origin == "https://github.com"
				},
				MaxAge: 3600,
			},
		),
	)
}
