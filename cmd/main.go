package main

import "github.com/justgast/google_calendar_on_kindle/pkg/use_cases"

const calendarId = "u80gdf74rouv41pl2a10r08ji0@group.calendar.google.com"

func main() {
	use_cases.GenerateCalendarImage(calendarId)
}
