package main

import (
	_ "pinyin/routers"

	"github.com/astaxie/beego"
)

func main() {
	// 设置日志级别
	logLevel, err := beego.GetConfig("String", "logLevel", "error")
	if err != nil {
		beego.Debug("get logLevel err!")
	}

	switch logLevel {
	case "error":
		beego.SetLevel(beego.LevelError)
	case "info":
		beego.SetLevel(beego.LevelInformational)
	case "debug":
		beego.SetLevel(beego.LevelDebug)
	default:
		beego.SetLevel(beego.LevelError)
	}

	beego.Run()
}
