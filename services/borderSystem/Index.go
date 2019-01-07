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
	file, err := GetOne(id, "_id")
	if nil != err {
		return err
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
	reply, _ := couchdb.Find(couchdb.Condition().Fields("_id", "name", "uploadTime", "size").
		Page(begin, limit))
	return reply
}

func GetList(begin, limit int) []interface{} {
	condition := couchdb.Condition().Append("status", "$eq", true).
		Fields("name", "uploadTime", "_id").
		Page(begin, limit)
	reply, _ := couchdb.Find(condition)
	return reply
}

func GetOne(key string, files ...string) (*fsFile, error) {
	condition := couchdb.Condition().Append("_id", "$eq", key).
		Append("status", "$eq", true)
	if 0 != len(files) {
		condition = condition.Fields(files...)
	}
	reply, _ := couchdb.Find(condition)
	if 0 != len(reply) {
		var fs fsFile
		couchdb.Decode(reply[0], &fs)
		return &fs, nil
	} else {
		return nil, &exceptions.Error{Msg: "no such this file", Code: 404}
	}
}