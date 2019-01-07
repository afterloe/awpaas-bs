package test

import (
	"testing"
	"encoding/json"
	"reflect"
	"../integrate/couchdb"
)

type fsFile struct {
	Id string `json:"id"`
	Name string `json:"name"`
	SavePath string `json:"savePath"`
	ContentType string `json:"contentType"`
	Key string `json:"key"`
	UploadTime int64 `json:"uploadTime"`
	Size int64 `json:"size"`
	Status bool `json:"status"`
}

func (this *fsFile) String() string {
	jsonBytes, _ := json.Marshal(this)
	return string(jsonBytes)
}

func Test_Reflect(t *testing.T) {
	item := map[string]interface{} {
		"ContentType": "image/png",
		"Key": "E598E47773DF409C9A3866731C80C5DF",
		"Name": "5bfce8384fe9e.png",
		"SavePath": "/tmp/filesystem",
		"Size": 2638496,
		"Status": true,
		"UploadTime": 1546850433,
		"_id": "c6224a3eb806bf144ea7da52e900211a",
		"_rev": "1-d29eed1f47e550e9c573b9873f4e1b61",
	}
	t.Log(item)
	t.Log(reflect.TypeOf(item))
	var fs fsFile
	err := couchdb.Decode(item, &fs)
	if err != nil {
		panic(err)
	}
	t.Log(&fs)
}