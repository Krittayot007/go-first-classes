package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *handlerImpl) HandleUserIdget(c *gin.Context, userId string) {

	// res, err := h.service.GuestLogin(c)
	// if err != nil {
	// 	errorUtil.CastErrorJson(c, err)
	// 	return
	// }

	c.JSON(200, map[string]interface{}{})

}
