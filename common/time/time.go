package time

import (
	"time"
)

const (
	Date     = "2006-01-02"
	DateTime = "2006-01-02 15:04:05"
)

func GetNowTime() string {
	now := time.Now()
	return now.Format(DateTime)
}
