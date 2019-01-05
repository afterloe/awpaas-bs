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

func GetList(skip, limit string) []interface{} {
	reply, _ :=couchdb.Read(dbName + "/_all_docs", map[string]interface{}{
		"skip": skip,
		"limit": limit,
		"include_docs": "true",
	})
	var list = make([]interface{}, 0)
	if "not_found" == reply["error"]{
		return list
	}
	for _, r := range (reply["rows"].([]interface{})) {
		doc := (r.(map[string]interface{}))["doc"].(map[string]interface{})
		delete(doc, "_rev")
		delete(doc, "SavePath")
		delete(doc, "Key")
		list = append(list, doc)
	}
	return list
}

func GetOne(key string) (map[string]interface{}, error) {
	condition := couchdb.Condition().Append("_id", "$eq", key).
		Append("Status", "$eq", true)
	reply, _ := couchdb.Find(dbName, condition)
	if 0 != len(reply) {
		return reply[0].(map[string]interface{}), nil
	} else {
		return nil, &exceptions.Error{Msg: "no such this file", Code: 404}
	}
}