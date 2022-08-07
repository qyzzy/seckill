package timeparser

import (
	"time"
)

func UnixToTime(data int64) time.Time {
	dataTime := time.Unix(data/1000, 0)
	return dataTime
}

func TimeToUnix(e time.Time) int64 {
	timeUnix, _ := time.Parse("2006-01-02 15:04:05", e.Format("2006-01-02 15:04:05"))
	return timeUnix.UnixNano() / 1e6
}
