package routers

import (
	"gvb_server/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdate)
	// 下面两个作废了
	router.GET("settings_email", settingsApi.SettingsEmailInfoView)
	router.PUT("settings_email", settingsApi.SettingsEmailInfoUpdate)
}
