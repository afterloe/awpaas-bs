package routers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"../util"
	"../config"
)

/**
	路由列表
 */
func Execute(route *gin.RouterGroup) {
	route.PUT("/file", FsUpload)
	route.GET("/file", FsList)
	route.DELETE("/file", FsDel)
	route.GET("/all/file", FsListAll)
	route.GET("/file/:key", FsFindOne)
	route.GET("/download/:key", FsDownload)
}

/**
	描述信息
 */
func Info(context *gin.Context) {
	info := config.Get("info").(map[string]interface{})
	context.SecureJSON(http.StatusOK, util.Success(info))
}

/**
	分页组件
 */
func pageCondition(context *gin.Context) (int, int) {
	begin, err := strconv.Atoi(context.DefaultQuery("bg", "0"))
	if nil != err {
		begin = 0
	}
	end, err := strconv.Atoi(context.DefaultQuery("ed", "10"))
	if nil != err {
		end = 10
	}
	limit := end - begin
	if 0 >= limit {
		limit = 10
	}
	return begin, limit
}
