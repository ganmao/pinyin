package libs

import (
	"github.com/astaxie/beego"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
)

// WX 全局的微信句柄
var WX *wechat.Wechat

func init(){
	cacheConf := beego.AppConfig.String("wxMemCache")
	MemCache := cache.NewMemcache(cacheConf)
	
	//配置微信参数
	wxConfig := &wechat.Config{
		AppID:          beego.AppConfig.String("wcAppId"),
		AppSecret:      beego.AppConfig.String("wcAppSecret"),
		Token:          beego.AppConfig.String("wcToken"),
		EncodingAESKey: beego.AppConfig.String("wcEncodingAESKey"),
		Cache: 			MemCache,
	}
	WX = wechat.NewWechat(wxConfig)

	AccessToken, err := WX.GetAccessToken()
	if err != nil {
		beego.Debug("Get AccessToken Error: ", err)
	}else{
		beego.Debug("Get AccessToken Sucess: ", AccessToken)
	}
}
