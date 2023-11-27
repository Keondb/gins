package models

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

type BannerModel struct {
	MODEL
	Path string `json:"path"` //图片路径
	//Sort int    `gorm:"size:4;default:0" json:"sort"` //排序字段
	Hash      string          `json:"hash"`                        //图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38" json:"name"`         //图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片类型，本地还还是七牛
	//ArticleModels []ArticleModel `gorm:"foreignKey:coverID" json:"-"`
	//AuthModels    []AuthModel    `gorm:"foreignKey:AvatarID" json:"-"`
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		// 本地图片
		err := os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
		}
	}
	return nil
}
