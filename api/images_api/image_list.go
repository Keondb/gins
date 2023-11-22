package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type Page struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

func (ImagesApi) ImageListView(c *gin.Context) {
	var imageList []models.BannerModel
	var cr Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 总数
	count := global.DB.Find(&imageList).RowsAffected
	if count > 0 {
		offset := cr.Limit * (cr.Page - 1)
		if offset <= 0 {
			offset = 0
		}
		global.DB.Limit(cr.Limit).Offset(offset).Order(cr.Sort).Find(&imageList)
	}

	res.OkWithData(gin.H{
		"total": count,
		"list":  imageList,
	}, c)
	return
}
