package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (SettingsApi) SettingsInfoUpdate(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return

	}

	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info
		//fmt.Println(info)
		//res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		var info config.Email
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = info
		//res.OkWithData(global.Config.Email, c)
	case "qq":
		var info config.QQ
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = info
		//res.OkWithData(global.Config.QQ, c)
	case "jwt":
		var info config.Jwt
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwt = info
		//res.OkWithData(global.Config.Jwt, c)
	case "qiniu":
		var info config.QiNiu
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = info
		//res.OkWithData(global.Config.QiNiu, c)
	default:
		res.FailWithMessage("没有对应配置信息", c)
		return
	}
	err = core.Setyaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)

}

//func (SettingsApi) SettingsInfoUpdate(c *gin.Context) {
//
//	var cr config.SiteInfo
//	err := c.ShouldBindJSON(&cr)
//	if err != nil {
//		res.FailWithCode(res.ArgumentError, c)
//	}
//
//	global.Config.SiteInfo = cr
//
//	err = core.Setyaml()
//	if err != nil {
//		global.Log.Error(err)
//		res.FailWithMessage(err.Error(), c)
//		return
//	}
//	res.OkWith(c)
//}
