package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type ImagesUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请输入id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名"`
}

func (ImagesApi) ImageUpdate(c *gin.Context) {
	var cr ImagesUpdateRequest
	// 检查参数是否符合要求
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		// 如果参数错误，返回错误信息
		res.FailWithError(err, &cr, c)
		return
	}
	// 更新数据库
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("文件不存在", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	//global.DB.Model(&imageModel).Updates(map[string]interface{}{"name": cr.Name})
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithMessage("更新成功", c)
	return
}
