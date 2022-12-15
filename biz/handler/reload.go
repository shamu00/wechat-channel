package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/shamu00/wechat-channel/util"
)

func Reload(_ context.Context, c *app.RequestContext) {
	util.MustInitGlobal()
	InitMessageHandler(util.GlobalWechatApiId, util.GlobalWechatApiSecret)
	c.JSON(consts.StatusOK, utils.H{
		"message": "success",
	})
}
