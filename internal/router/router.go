package router

import (
	"github.com/SalhSThai/go-first-classes/internal/handler"

	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.Engine, h handler.Handler) {

	authRoute := r.Group("/auth")
	{
		authRoute.POST("/login", h.HandleLoginPost)
		authRoute.POST("/register", h.HandleRegisterPost)
		authRoute.POST("/remember", h.HandleRememberPost)

	}
	userRoute := r.Group("/user")
	{
		userRoute.GET("/", h.HandleUserAllGet)
		userRoute.POST("/", h.HandleUserCreatePost)
		userRoute.GET("/:user-id", func(ctx *gin.Context) { h.HandleUserIdget(ctx, ctx.Param("user-id")) })
		userRoute.PUT("/:user-id", func(ctx *gin.Context) { h.HandleUserUpdatePut(ctx, ctx.Param("user-id")) })

	}
}
