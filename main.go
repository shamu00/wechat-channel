// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/shamu00/wechat-channel/biz/handler"
	"github.com/shamu00/wechat-channel/util"
)

func main() {
	util.MustInitGlobal()
	handler.InitMessageHandler(util.GlobalWechatApiId, util.GlobalWechatApiSecret)
	h := server.Default(config.Option{F: func(o *config.Options) {
		o.Addr = ":80"
	}})

	register(h)
	h.Spin()
}
