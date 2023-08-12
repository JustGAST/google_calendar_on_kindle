package calendar_utils

import "time"

func GetStartDay() time.Time {
	t := time.Now().AddDate(0, 0, -time.Now().Day()+1)
	return t.AddDate(0, 0, -int(t.Weekday())+1)
}
