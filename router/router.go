package router

import (
	"pert1/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.RootHandler)

	r.POST("/post", handler.PostHandler)

	r.GET("/jsononline", handler.RootHandlerJsonOnline)
}
