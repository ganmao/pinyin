package routers

import (
	"pinyin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.WeChatController{})
}
