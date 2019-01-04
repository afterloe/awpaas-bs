package borderSystem

import (
	"fmt"
	"time"
)

type fsFile struct {
	name, savePath, contentType, key string
	uploadTime, size int64
	status bool
}

func (this *fsFile) String() string {
	return fmt.Sprintf("name: %s savePaht: %s contentType: %s key: %s, uploadTime: %s, size: %d, status %v",
		this.name, this.savePath, this.contentType, this.key, time.Unix(this.uploadTime, 0).Format(timeFormat),
		this.size, this.status)
}
