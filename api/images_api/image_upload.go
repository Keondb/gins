package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/utils"
	"io/fs"
	"os"
	"path"
	"strings"
)

type FileUploadResponse struct {
	FileName  string `json:"fileName"`   // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// '.xbm','.tif','pjp','.svgz','jpg','jpeg','ico','tiff','.gif','svg','.jfif','.webp','.png','.bmp','pjpeg','.avif'
var (
	// 图片上传的白名单
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
	}
)

func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在相对应的文件", c)
		return
	}
	basePath := global.Config.Upload.Path
	_, err = os.ReadFile(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
		fmt.Println(err)
	}

	var resList []FileUploadResponse

	for _, file := range fileList {
		fileName := file.Filename
		nameList := strings.Split(fileName, ".")
		suffix := strings.ToLower(nameList[len(nameList)-1])
		if !utils.InList(suffix, WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue
		}

		filePath := path.Join(basePath, fileName)
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过设定大小。当前大小为：%.2fMB,设定大小为：%dMB", size, global.Config.Upload.Size),
			})
			continue
		}
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			//global.Log.Error(err)
			continue
		}
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "图片上传成功",
		})
	}
	res.OkWithData(resList, c)
	//fileHeaderc, err := c.FormFile("image")
	//if err != nil {
	//	res.FailWithMessage(err.Error(), c)
	//}

}
