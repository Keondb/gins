package models

import "gvb_server/models/ctype"

type UserModel struct {
	MODEL
	NickName       string           `gorm:"size:36" json:"nick_name"`                                                                //昵称
	UserName       string           `gorm:"size:36" json:"user_name"`                                                                //用户名
	Password       string           `gorm:"size:128" json:"password"`                                                                //密码
	Avatar         string           `gorm:"size:256" json:"avatar_id"`                                                               //头像
	Email          string           `gorm:"size:128" json:"email"`                                                                   //邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                                                                      //手机号
	Addr           string           `gorm:"size:64" json:"addr"`                                                                     //地址
	Token          string           `gorm:"size:64" json:"token"`                                                                    //其他平台的唯一id
	IP             string           `gorm:"size:20" json:"ip"`                                                                       //ip地址
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"`                                                            // 权限 1管理 2普通 3游客
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`                                                     // 注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignKey:UserID" json:"-"`                                                              //发布的文章列表
	CollectsModels []ArticleModel   `gorm:"many2many:user_collect_models; joinForeignKey :UserID;JoinReferences:ArticleID" json:"-"` //收藏的文章列表
}

//type AuthModel struct {
//	MODEL
//	NickName       string                        `gorm:"size:42" json:"nick_name"`
//	UserName       string                        `gorm:"size:36" json:"user_name"`
//	Password       string                        `gorm:"size:64" json:"password"`
//	AvatarId       uint                          `json:"avatar_id"`
//	Avatar         ImageMole                     `json:"-"`
//	Email          string                        `json:"email"`
//	Tel            string                        `gorm:"size:18" json:"tel"`
//	Addr           string                        `gorm:"size:64" json:"addr"`
//	Token          string                        `gorm:"size:64" json:"token"`
//	IP             string                        `gorm:"size:20" json:"IP"`
//	Role           int                           `gorm:"size:4;default:1" json:"role"`                                                       // 权限 1管理 2普通 3游客
//	SignStatus     status_type.AccountStatusType `gorm:"type=smallint(6)" json:"sign_status"`                                                // 注册来源
//	ArticleModels  []ArticleModel                `gorm:"foreignKey:AuthId" json:"-"`                                                         //发布的文章列表
//	CollectsModels []ArticleModel                `gorm:"many2many:auth2_collects; joinForeignKey :AuthID;JoinReferences:ArticleID" json:"-"` //收藏的文章列表
//	SiteModels     []SiteModel                   `gorm:"many2many:auth_sites" json:"-"`                                                      //收藏的网站列表
//}
