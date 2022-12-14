// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "github.com/shamu00/wechat-channel/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.GET("/reload", handler.Reload)
	r.GET("/message", handler.Message)
	// your code ...
}
