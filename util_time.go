package util

import (
	"time"

	"github.com/gogf/gf/util/gconv"
)

// Msec timestamp format to Y-m-d H:i:s.
func MsecFormat(msectime int64) string {
	msectimeUnix := gconv.Int64(msectime / 1000)
	timeFormat := time.Unix(msectimeUnix, 0).Format("2006-01-02 15:04:05")
	return timeFormat
}
