package test

import (
	"testing"
	"../services/borderSystem"
	"encoding/json"
	"reflect"
	"fmt"
)

func Test_utils(t *testing.T) {
	fs := borderSystem.Default("yyy.png", "file", 8 >> 12)
	t.Log(fs)
	value := reflect.ValueOf(fs)
	fmt.Println(value.FieldByName("name"))
	//ty := reflect.TypeOf(fs)
	//for i := 0; i < value.NumField(); i++ {
	//	v := value.Field(i)
	//	if v.IsValid() {
	//		t.Log(fmt.Sprintf("%v -> %v", ty.Field(i).Name, v.Interface()))
	//	}
	//}

	buf, err := json.Marshal(*fs)
	if nil != err {
		t.Error(err)
	}
	t.Log(string(buf))
}
