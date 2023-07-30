package models

type BannerModel struct {
	MODEL
	Path string `json:"path"` //图片路径
	//Sort int    `gorm:"size:4;default:0" json:"sort"` //排序字段
	Hash string `json:"hash"`                //图片的hash值，用于判断重复图片
	Name string `gorm:"size:38" json:"name"` //图片名称
	//ArticleModels []ArticleModel `gorm:"foreignKey:coverID" json:"-"`
	//AuthModels    []AuthModel    `gorm:"foreignKey:AvatarID" json:"-"`
}
