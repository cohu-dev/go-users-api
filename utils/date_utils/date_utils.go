package date_utils

import "time"

const apiDateLayout = "2006-01-02T15:04:05.000Z"

func GetTime()time.Time{
	return time.Now().UTC()
}

func GetNowString() string {
	return  GetTime().Format(apiDateLayout)
}
