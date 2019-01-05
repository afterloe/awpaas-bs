package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../util"
	"../services/borderSystem"
	"strconv"
)

/**
	文件上传
 */
func FsUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if nil != err {
		ctx.JSON(http.StatusBadRequest, util.Fail(400, "file not found."))
		return
	}
	fs := borderSystem.Default(file.Filename, file.Header.Get("Content-Type"), file.Size)
	object, err := fs.SaveToDB()
	if nil != err {
		ctx.JSON(http.StatusInternalServerError, util.Error(err))
		return
	}
	err = ctx.SaveUploadedFile(file, fs.GeneratorSavePath())
	if nil != err {
		ctx.JSON(http.StatusInternalServerError, util.Fail(500, "io exception."))
		return
	}
	ctx.JSON(http.StatusOK, util.Success(object))
}

func FsListAll(ctx *gin.Context) {
	begin, limit := pageCondition(ctx)
	reply := borderSystem.GetAll(strconv.Itoa(begin), strconv.Itoa(limit))
	ctx.JSON(http.StatusOK, util.Success(reply))
}

func FsList(ctx *gin.Context) {
	begin, limit := pageCondition(ctx)
	reply := borderSystem.GetList(begin, limit)
	ctx.JSON(http.StatusOK, util.Success(reply))
}

func FsFindOne(ctx *gin.Context) {
	key := ctx.Param("key")
	if "" == key {
		ctx.JSON(http.StatusBadRequest, util.Fail(400, "key not found"))
		return
	}
	reply, err := borderSystem.GetOne(key)
	if nil != err {
		ctx.JSON(http.StatusInternalServerError, util.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, util.Success(reply))
}