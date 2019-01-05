package couchdb

import "../../util"

type condition map[string]interface{}

func (this *condition) String() string {
	reply, _ := util.FormatToString(this)
	return reply
}

func Condition() condition {
	return condition{}
}