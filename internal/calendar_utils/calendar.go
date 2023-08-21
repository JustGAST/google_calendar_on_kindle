package calendar_utils

import "time"

func GetStartDay() time.Time {
	now := time.Now()
	// Get first day of month
	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).
		AddDate(0, 0, -time.Now().Day()+1)

	// Get first monday before first day of month
	return t.AddDate(0, 0, -int(t.Weekday())+1)
}
