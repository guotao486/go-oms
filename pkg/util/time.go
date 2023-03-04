package util

import "time"

func UnixToString(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return FormatYMDHIS(tm)
}

func FormatYMDHIS(tm time.Time) string {
	return tm.Format("2006-01-02 03:04:05")
}

func FormatYMD(tm time.Time) string {
	return tm.Format("2006-01-02")
}

func FormatHIS(tm time.Time) string {
	return tm.Format("03:04:05")
}
