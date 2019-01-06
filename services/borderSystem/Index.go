package borderSystem

import (
	"../../integrate/couchdb"
	"../../exceptions"
	"../../util"
	"../../config"
	"fmt"
	"time"
)

var (
	root string
)

func init() {
	cfg := config.Get("custom")
	root = config.GetByTarget(cfg, "root").(string)
}

func (this *fsFile) SaveToDB() (map[string]interface{}, error){
	return couchdb.Create(this)
}

func (this *fsFile) Del(f ...bool) error {
	if 0 != len(f) { // 强制删除

	} else { // 逻辑删除

	}

	return nil
}
/**
	TODO
*/
func Del(id string, f ...bool) (error) {
	reply, err := GetOne(id, "_id", "Key", "SavePath", "Status")
	if nil != err {
		return err
	}
	file := &fsFile{
		id: reply["_id"].(string),
		Key: reply["Key"].(string),
		SavePath: reply["SavePath"].(string),
		UploadTime: reply["UploadTime"].(int64),
		Status: reply["Status"].(bool),
	}
	return file.Del(f...)
}

func Default(name, contentType string, size int64) *fsFile {
	return &fsFile{
		SavePath: root,
		Key: util.GeneratorUUID(),
		UploadTime: time.Now().Unix(),
		Name: name,
		ContentType: contentType,
		Size: size,
		Status: true,
	}
}

func (this *fsFile) GeneratorSavePath() string {
	return fmt.Sprintf("%s/%s", this.SavePath, this.Key)
}

func GetAll(begin, limit int) []interface{} {
	reply, _ := couchdb.Find(couchdb.Condition().Fields("_id", "Name", "UploadTime", "Size").
		Page(begin, limit))
	return reply
}

func GetList(begin, limit int) []interface{} {
	condition := couchdb.Condition().Append("Status", "$eq", true).
		Fields("Name", "UploadTime", "_id").
		Page(begin, limit)
	reply, _ := couchdb.Find(condition)
	return reply
}

func GetOne(key string, files ...string) (map[string]interface{}, error) {
	condition := couchdb.Condition().Append("_id", "$eq", key).
		Append("Status", "$eq", true)
	if 0 != len(files) {
		condition = condition.Fields(files...)
	}
	reply, _ := couchdb.Find(condition)
	if 0 != len(reply) {
		return reply[0].(map[string]interface{}), nil
	} else {
		return nil, &exceptions.Error{Msg: "no such this file", Code: 404}
	}
}