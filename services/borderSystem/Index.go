package borderSystem

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../../util"
	"../../config"
	"time"
)

var (
	root, dbName string
)

func init() {
	dbName = "file-system"
	custom := config.Get("custom")
	rootCfg := config.GetByTarget(custom, "root")
	if nil != rootCfg {
		root = "/tmp/filesystem"
	} else {
		root = rootCfg.(string)
	}
}

/**
	文件上传
 */
func FsUpload(context *gin.Context) {
	file, err := context.FormFile("file")
	if nil != err {
		context.JSON(http.StatusBadRequest, util.Fail(400, "file not found."))
		return
	}
	fs := &fsFile{
		name: file.Filename,
		savePath: root,
		contentType: file.Header.Get("Content-Type"),
		key: util.GeneratorUUID(),
		uploadTime: time.Now().Unix(),
		size: file.Size,
		status: true,
	}
	object, err := saveToCouchDB(fs.generatorMap())
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