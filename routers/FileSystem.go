package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../util"
	"../services/borderSystem"
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
	fs := borderSystem.Default(file.Filename, file.Header.Get("Content-Type"), file.Size)
	object, err := fs.SaveToDB()
	if nil != err {
		context.JSON(http.StatusInternalServerError, util.Error(err))
		returni
	}
	err = context.SaveUploadedFile(file, fs.GeneratorSavePath())
	if nil != err {
		context.JSON(http.StatusInternalServerError, util.Fail(500, "io exception."))
		return
	}
	context.JSON(http.StatusOK, util.Success(object))
}