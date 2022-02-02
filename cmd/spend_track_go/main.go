package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	var port string = getEnvVariable("SERVER_PORT")
	var portString string = fmt.Sprintf(":%s", port)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "index"})
	})

	log.Print(fmt.Sprintf("Running on port: %s", portString))
	r.Run(portString)
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Could not load .env file")
	}

	return os.Getenv(key)
}
