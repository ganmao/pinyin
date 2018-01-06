package controllers

import (
	"github.com/astaxie/beego"
)

type WeChatController struct {
	beego.Controller
}

func (c *WeChatController) Get() {
	c.Data["Website"] = "pinyin"
	c.Data["Email"] = "zdl0812@163.com"
	c.TplName = "index.tpl"
}
