package utils

import "time"

const layout = "2006-01-02 15:04:05"

func TimeToString(t time.Time) string {
	return t.Format(layout)
}

func StringToTime(str string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}
