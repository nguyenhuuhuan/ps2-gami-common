package utils

import "time"

func ConvertTimeToMiliSeconds(time time.Time) int64 {
	return time.UnixNano() / 1000000
}
