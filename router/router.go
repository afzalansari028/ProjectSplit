package router

import (
	"splitbill/handler"

	"github.com/gin-gonic/gin"
)

func EndPoints(r *gin.Engine) {

	r.GET("/test", handler.TestAPI)

	r.POST("/split", handler.DoSplit)
}
