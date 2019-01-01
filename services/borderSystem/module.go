package borderSystem

import (
	"fmt"
	"time"
	"../../integrate/couchdb"
)

var (
	timeFormat string
)

func init() {
	timeFormat = "2006-01-02 - 15:04:05"
}

type fsFile struct {
	name, savePath, contentType, key string
	uploadTime, size int64
	status bool
}

func (this *fsFile) generatorSavePath() string {
	return fmt.Sprintf("%s/%s", this.savePath, this.key)
}

func (this *fsFile) String() string {
	return fmt.Sprintf("name: %s savePaht: %s contentType: %s key: %s, uploadTime: %s, size: %d, status %v",
		this.name, this.savePath, this.contentType, this.key, time.Unix(this.uploadTime, 0).Format(timeFormat),
		this.size, this.status)
}

func (this *fsFile) generatorMap() map[string]interface{} {
	return map[string]interface{}{
		"name": this.name,
		"savePath": this.savePath,
		"contentType": this.contentType,
		"key": this.key,
		"uploadTime": this.uploadTime,
		"size": this.size,
		"status": this.status,
	}
}

func (this *fsFile) saveToCouchDB() (map[string]interface{}, error){
	return couchdb.Create(dbName, *this)
}