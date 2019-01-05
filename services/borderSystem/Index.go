package borderSystem

import (
	"../../integrate/couchdb"
	"../../util"
	"../../config"
	"fmt"
	"time"
)

var (
	root, dbName string
)

func init() {
	cfg := config.Get("custom")
	root = config.GetByTarget(cfg, "root").(string)
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

func (this *fsFile) GeneratorSavePath() string {
	return fmt.Sprintf("%s/%s", this.savePath, this.key)
}