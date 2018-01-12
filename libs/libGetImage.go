package libs

import (
	"os"
	"io"
    "io/ioutil"
	"net/http"
	"bytes"

	"github.com/astaxie/beego"
)

// 下载微信图片，保存到指定目录
func GetWeChatImage(name string, url string, localPath string) (n int64, err error){
	beego.Debug("Download Image name : ", name)
	beego.Debug("Download Image url : ", url)
	beego.Debug("Download Image localPath : ", localPath)
	
	out, err := os.Create(localPath + "/" + name)
	defer out.Close()
	if err != nil {
		beego.Debug("Create Image Err : ", err)
	}
	
    resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		beego.Debug("Get Image Err : ", err)
	}
	
	pix, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug("ReadAll Image Err : ", err)
	}

	n, err = io.Copy(out, bytes.NewReader(pix))
	if err != nil {
		beego.Debug("Copy Image Err : ", err)
	}

	return
}