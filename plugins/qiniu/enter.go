package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"gvb_server/config"
	"gvb_server/global"
	"time"
)

// getToken  获取返上传的token
func getToken(niu config.QiNiu) string {
	accessKey := niu.AccessKey
	secretKey := niu.SecretKey
	bucket := niu.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// getCfg  获取上传的配置
func getCfg(niu config.QiNiu) storage.Config {
	cfg := storage.Config{}
	// 空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(niu.Zone))
	cfg.Zone = &zone
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用cdn上传加密
	cfg.UseCdnDomains = false
	return cfg
}

// UploadImage 上传图片 文件数组，前缀
func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
	if !global.Config.QiNiu.Enable {
		return "", errors.New("请启用七牛云上传")
	}
	// 文件名不能重复
	niu := global.Config.QiNiu
	if niu.SecretKey == "" || niu.AccessKey == "" {
		return "", errors.New("请配置AccessKey及SecretKey")
	}
	if float64(len(data))/1024/1024 > niu.Size {
		return "", errors.New("文件超过设定大小")
	}
	upToken := getToken(niu)
	cfg := getCfg(niu)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	// 获取当前时间
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s__%s", prefix, now, imageName)
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", niu.CDN, ret.Key), nil
}
