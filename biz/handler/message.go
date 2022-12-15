package handler

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/shamu00/wechat-channel/util"
	sdk "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	oaconfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

type MessageResponse struct {
	Text string    `xml:"text"`
	Time time.Time `xml:"time"`
}

var (
	wechat *sdk.Wechat
	mem    *cache.Memory
	cfg    *oaconfig.Config
	client *officialaccount.OfficialAccount
)

const EncodingKey = "WECHAT_ENCODING_KEY"
const Token = "WECHAT_CHANNEL_TOKEN"

func InitMessageHandler(appId, appSecret string) {
	wechat = sdk.NewWechat()
	mem = cache.NewMemory()
	encodingKey := os.Getenv(EncodingKey)
	token := os.Getenv(Token)
	log.Printf("[Info]EncodingKey length:%d, Token length:%d", len(encodingKey), len(token))
	cfg = &oaconfig.Config{
		AppID:          appId,
		AppSecret:      appSecret,
		Token:          token,
		EncodingAESKey: encodingKey,
		Cache:          mem,
	}
	client = wechat.GetOfficialAccount(cfg)
}

func Message(ctx context.Context, c *app.RequestContext) {
	request, err := adaptor.GetCompatRequest(c.GetRequest())
	if err != nil {
		fmt.Printf("[Error]GetCompatRequest, err:%v, origin:%v", err, c.GetRequest())
		return
	}
	server := client.GetServer(request, adaptor.GetCompatResponseWriter(&c.Response))
	h := &handler{ctx}
	server.SetMessageHandler(h.messageHandler)
	err = server.Serve()
	if err != nil {
		fmt.Printf("[Error]serve:%v", err)
		return
	}
	// 发送回复的消息
	err = util.Retry(3, 0, func() error {
		return server.Send()
	})
	if err != nil {
		log.Printf("[Error]send reply message:%v", err)
	}
}

type handler struct {
	context.Context
}

func (h handler) messageHandler(msg *message.MixMessage) *message.Reply {
	log.Printf("[Debug]message:%+v", msg)
	req := util.NewDefaultCompletionRequest(msg.Content, msg.CommonToken.GetOpenID())
	var err error
	var res *gogpt.CompletionResponse
	err = util.Retry(3, 500*time.Millisecond, func() error {
		var e error
		res, e = util.GlobalChatGptClient.Talk(h.Context, req)
		return e
	})
	if err != nil {
		fmt.Printf("[Error]calling chatgpt client, from:%v, input:%v, error:%v",
			msg.CommonToken.GetOpenID(), msg.Content, err)
		return &message.Reply{
			MsgType: message.MsgTypeText,
			MsgData: message.NewText(fmt.Sprintf("内部错误，请尝试重试 :)")),
		}
	}
	return &message.Reply{
		MsgType: message.MsgTypeText,
		MsgData: res.Choices[0].Text,
	}
}
