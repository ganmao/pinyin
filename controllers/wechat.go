package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

type WeChatController struct {
	beego.Controller
}

func (c *WeChatController) Any() {
	beego.Debug("Begin Any" )
	//配置微信参数
	config := &wechat.Config{
		AppID:          beego.AppConfig.String("wcAppId"),
		AppSecret:      beego.AppConfig.String("wcAppSecret"),
		Token:          beego.AppConfig.String("wcToken"),
		EncodingAESKey: beego.AppConfig.String("wcEncodingAESKey"),
	}
	wc := wechat.NewWechat(config)

	beego.Debug("Get data:", c.Ctx.Input.Method(), c.Ctx.Input.Site())

	// 传入request和responseWriter
	server := wc.GetServer(c.Ctx.Request, c.Ctx.ResponseWriter)

	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{message.MsgTypeText, text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}
