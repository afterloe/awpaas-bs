package borderSystem

import (
	"fmt"
	"time"
)

var (
	timeFormat string
)

func init() {
	timeFormat = "2006-01-02 - 15:04:05"
}

type fsFile struct {
	Name, SavePath, ContentType, Key string
	UploadTime, Size int64
	Status bool
}

func (this *fsFile) String() string {
	return fmt.Sprintf("name: %s savePaht: %s contentType: %s key: %s, uploadTime: %s, size: %d, status %v",
		this.Name, this.SavePath, this.ContentType, this.Key, time.Unix(this.UploadTime, 0).Format(timeFormat),
		this.Size, this.Status)
}
