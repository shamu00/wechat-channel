package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"time"
)

type MessageResponse struct {
	Text string    `xml:"text"`
	Time time.Time `xml:"time"`
}

func Message(ctx context.Context, c *app.RequestContext) {
	c.XML(consts.StatusOK, &MessageResponse{
		Text: "hello world",
		Time: time.Now(),
	})
}
