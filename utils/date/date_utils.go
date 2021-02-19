package date

import "time"

const (
	apiDateLayout = "2020-19-02T23:27:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
