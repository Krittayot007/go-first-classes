package handler

import (
	"github.com/SalhSThai/go-first-classes/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	HandleLoginPost(c *gin.Context)
	HandleRegisterPost(c *gin.Context)
	HandleRememberPost(c *gin.Context)
	HandleUserAllGet(c *gin.Context)
	HandleUserCreatePost(c *gin.Context)
	HandleUserIdget(c *gin.Context, userId string)
	HandleUserUpdatePut(c *gin.Context, userId string)
}

type handlerImpl struct {
	service service.Service
}

func New(s service.Service) (Handler, error) {
	return &handlerImpl{
		service: s,
	}, nil
}
