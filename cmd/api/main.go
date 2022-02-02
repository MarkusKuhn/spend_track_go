package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

import "../../pkg/api/transactions"

func main() {
	router := gin.Default()
	var port string = getEnvVariable("SERVER_PORT")
	var portString string = fmt.Sprintf(":%s", port)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "index"})
	})

	router.GET("/transactions", func(c *gin.Context) {
		c.JSON(200, gin.H{"transactions": transactions.GetTransactions()})
	})

	log.Print(fmt.Sprintf("Running on port: %s", portString))
	router.Run(portString)
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Could not load .env file")
	}

	return os.Getenv(key)
}
