package controllers

import "github.com/gin-gonic/gin"

type VideoController struct {}

var videoController = &VideoController{}

func (vc *VideoController) PostVideoMetaData(c *gin.Context) {
	var videoMeta models.Video
}