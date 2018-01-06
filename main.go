package main

import (
	"fmt"
	// _ "pinyin/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

func hello(ctx *context.Context) {
	//配置微信参数
	config := &wechat.Config{
		AppID:          "wxe5a524a81ebaded6",
		AppSecret:      "57206d630acdab6f0b4176b4f3743ade",
		Token:          "zhanghaolin1017",
		EncodingAESKey: "7bT4Et5UdajhtC0vRn0Ca4Chf8rwBhYDbRlC6UjSABK",
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(ctx.Request, ctx.ResponseWriter)

    beego.Debug("get data:", ctx.Request)
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

func main() {
	beego.Any("/weixin", hello)
	beego.Run()
}
