package borderSystem

import (
	"../../integrate/couchdb"
	"../../util"
	"fmt"
	"time"
)

var (
	root, timeFormat string
	host, dbName string
)

func init() {
	root = "/tmp/filesystem"
	timeFormat = "2006-01-02 - 15:04:05"
	host = "mine:5984"
	dbName = "file-system"
}

func (this *fsFile) SaveToDB() (map[string]interface{}, error){
	return couchdb.Create(dbName, this)
}

func Default(name, contentType string, size int64) *fsFile {
	return &fsFile{
		savePath: root,
		key: util.GeneratorUUID(),
		uploadTime: time.Now().Unix(),
		name: name,
		contentType: contentType,
		size: size,
		status: true,
	}
}

func GetRoot() string {
	return root
}

//func (this *fsFile) GeneratorMap() map[string]interface{} {
//	return map[string]interface{}{
//		"name": this.Name,
//		"savePath": this.SavePath,
//		"contentType": this.ContentType,
//		"key": this.Key,
//		"uploadTime": this.UploadTime,
//		"size": this.Size,
//		"status": this.Status,
//	}
//}

func (this *fsFile) GeneratorFullPath() string {
	return fmt.Sprintf("%s/%s", this.savePath, this.key)
}