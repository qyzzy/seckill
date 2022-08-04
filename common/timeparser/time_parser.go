package timeparser

import (
	"time"
)

func UnixToTime(data int64) (dataTime time.Time) {
	dataTime = time.Unix(data/1000, 0)
	return
}
