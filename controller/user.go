package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/model/response"
)

func init() {
	core.Router.GET("/any", func(context *gin.Context) {
		context.JSON(200, response.Response{StatusCode: 200, StatusMsg: "ok"})
	})
}