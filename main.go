package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Router
	router.GET("/", HelloWorld)

	fmt.Println("Server running on http://localhost:2008")
	router.Run(":2008")
}

func HelloWorld(c *gin.Context) {
	message := gin.H{"message": "Hello, World!"}

	c.JSON(http.StatusOK, message)
}
