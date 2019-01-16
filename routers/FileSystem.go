package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../util"
	"../services/borderSystem"
	"os"
)

/**
	文件上传
 */
func FsUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if nil != err {
		ctx.SecureJSON(http.StatusBadRequest, util.Fail(400, "file not found."))
		return
	}
	fs := borderSystem.Default(file.Filename, file.Header.Get("Content-Type"), file.Size)
	object, err := fs.SaveToDB()
	if nil != err {
		ctx.SecureJSON(http.StatusInternalServerError, util.Error(err))
		return
	}
	err = ctx.SaveUploadedFile(file, fs.GeneratorSavePath())
	if nil != err {
		fs.Del(true)
		ctx.SecureJSON(http.StatusInternalServerError, util.Fail(500, "io exception."))
		return
	}
	ctx.SecureJSON(http.StatusOK, util.Success(object))
}

func FsListAll(ctx *gin.Context) {
	reply := borderSystem.GetAll(pageCondition(ctx))
	ctx.SecureJSON(http.StatusOK, util.Success(reply))
}

func FsList(ctx *gin.Context) {
	reply := borderSystem.GetList(pageCondition(ctx))
	ctx.SecureJSON(http.StatusOK, util.Success(reply))
}

func FsDel(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	f := ctx.DefaultQuery("f", "")
	if "" == id {
		ctx.SecureJSON(http.StatusBadRequest, util.Fail(400, "lack param -> id"))
		return
	}
	var (
		err error
	)
	if "" == f {
		err = borderSystem.Del(id)
	} else {
		err = borderSystem.Del(id, true)
	}
	if nil != err {
		ctx.SecureJSON(http.StatusInternalServerError, util.Error(err))
		return
	}
	ctx.SecureJSON(http.StatusOK, util.Success("done"))
}

func FsFindOne(ctx *gin.Context) {
	key := ctx.Param("key")
	if "" == key {
		ctx.SecureJSON(http.StatusBadRequest, util.Fail(400, "key not found"))
		return
	}
	reply, err := borderSystem.GetOne(key)
	if nil != err {
		ctx.SecureJSON(http.StatusInternalServerError, util.Error(err))
		return
	}
	ctx.SecureJSON(http.StatusOK, util.Success(*reply))
}

func FsDownload(ctx *gin.Context) {
	key := ctx.Param("key")
	if 32 > len(key) {
		ctx.SecureJSON(http.StatusBadRequest, util.Fail(400, "参数错误"))
		return
	}
	file, err := borderSystem.GetOne(key)
	if nil != err {
		ctx.SecureJSON(http.StatusInternalServerError, util.Error(err))
		return
	}
	stat, err := os.Stat(file.GeneratorSavePath())
	if nil != err {
		ctx.SecureJSON(http.StatusNotFound, util.Fail(404, "file is not found"))
		return
	}
	ctx.Status(200)
	ctx.Header("Content-Type", file.ContentType)
	ctx.Header("Content-Disposition", "attachment;filename=" + string([]byte(file.Name)))
	ctx.Header("Content-Length", string(stat.Size()))
	ctx.File(file.GeneratorSavePath())
}