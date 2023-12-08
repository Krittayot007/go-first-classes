package main

import (
	"fmt"

	"github.com/SalhSThai/go-first-classes/src/model"
	"github.com/SalhSThai/go-first-classes/src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"
	engine := gin.Default()
	r := model.Server{Engine: engine}

	router.UserRouter(r, "/user")
	fmt.Println("Server started Port: ", port)
	r.Engine.Run(":8080")
}
