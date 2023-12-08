package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"
	engine := gin.Default()

	engine.POST("/login", LoginHandler)

	fmt.Println("Server started Port: ", port)
	engine.Run(":8080")
}
func LoginHandler(c *gin.Context) {

	c.JSON(200, c.Request.Body)
}
