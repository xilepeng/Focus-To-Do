package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "success") })
	}

	return r
}
