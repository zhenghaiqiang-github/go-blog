package api

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	_ "github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	_ "github.com/qiniu/go-sdk/v7/storage"
	"go-blob/common"
	"go-blob/config"
	"net/http"
)

func (*Api) QiniuToken(w http.ResponseWriter, r *http.Request) {
	//自定义凭证有效期(2小时,Expires单位为秒，为上传凭证的有效时间)
	bucket := "mszlu"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(config.Cfg.System.QiniuAccessKey,
		config.Cfg.System.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	common.Success(w, upToken)
}
