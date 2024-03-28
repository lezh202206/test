package templates

const TransportHttp = `
package transport_http

import (
	"gitee.com/zhubaoe-go/jordan/auth"
	"gitee.com/zhubaoe-go/jordan/go-kit/transport"
	chttp "gitee.com/zhubaoe-go/jordan/http"
	"github.com/gin-gonic/gin"
	kitTransport "github.com/go-kit/kit/transport/http"

	"{{.ImportPath -}} /svc"
)

var endpoints svc.Endpoints
var options []kitTransport.ServerOption

// http路由处理
func MakeHTTPHandler(_endpoints svc.Endpoints, _options ...kitTransport.ServerOption) *gin.Engine {
	endpoints = _endpoints

	serverOptions := transport.DefaultServerOptions()
	serverOptions = append(serverOptions, _options...)
	options = serverOptions

	router := chttp.NewGin()
	// 添加权限控制
	router.Use(chttp.UseAuth(auth.ToPermissions(Permissions)))

	// 绑定路由
	//defaultRouter(router)

	return router
}
`
