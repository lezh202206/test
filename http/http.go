package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/t_viper"
	"test/tracing"
)

func NewGin() *gin.Engine {
	gin.SetMode(t_viper.HttpViper().Mode)
	router := gin.New()
	router.Use(
		tracing.Http().Middleware(),
	)

	// 添加健康检查路由;
	router.GET("/Health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})
	// 网关的负载均衡检查
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})

	return router
}

// TODO
func Router() {

}
