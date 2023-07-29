package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.FailWithCode(res.SettingsError, c)
	//c.JSON(200, gin.H{"msg": "xxx11"})
}
