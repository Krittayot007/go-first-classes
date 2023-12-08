package router

import (
	"github.com/SalhSThai/go-first-classes/src/controller"
	"github.com/SalhSThai/go-first-classes/src/model"
)

func UserRouter(server model.Server, PATH string) {
	server.Engine.GET(PATH, controller.GetUser)
}
