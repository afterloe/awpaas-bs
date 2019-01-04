package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../util"
	"../services/borderSystem"
	"time"
)

/**
	文件上传
 */
func FsUpload(context *gin.Context) {
	file, err := context.FormFile("file")
	if nil != err {
		context.JSON(http.StatusBadRequest, util.Fail(400, "file not found."))
		return
	}
	fs := &borderSystem.FsFile{
		Name: file.Filename,
		SavePath: borderSystem.GetRoot(),
		ContentType: file.Header.Get("Content-Type"),
		Key: util.GeneratorUUID(),
		UploadTime: time.Now().Unix(),
		Size: file.Size,
		Status: true,
	}
	object, err := fs.SaveToDB()
	if nil != err {
		context.JSON(http.StatusInternalServerError, util.Error(err))
		return
	}
	err = context.SaveUploadedFile(file, fs.generatorSavePath())
	if nil != err {
		context.JSON(http.StatusInternalServerError, util.Fail(500, "io exception."))
		return
	}
	context.JSON(http.StatusOK, util.Success(object))
}