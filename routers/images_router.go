package routers

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	ImagesApi := api.ApiGroupApp.ImagesApi
	router.POST("images", ImagesApi.ImageUploadView)
	router.GET("images", ImagesApi.ImageListView)
}
