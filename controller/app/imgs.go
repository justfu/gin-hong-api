package app

import (
	"fmt"
	"gin/config"
	"gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetImgs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"imgPath": service.GetImg(),
	})
}

func UploadFileTest(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		panic(err)
	}

	fmt.Println(config.GetProjectTruePath() + "/upload/" + header.Filename)

	c.SaveUploadedFile(header, config.GetProjectTruePath()+"/upload/"+header.Filename)
	service.UploadFile(header)
}
