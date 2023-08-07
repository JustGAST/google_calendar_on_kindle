package google_api

import (
	"log"

	"google.golang.org/api/calendar/v3"
)

func GetCalendars(srv *calendar.Service) {
	calendars, err := srv.CalendarList.List().Do()
	if err != nil {
		log.Fatalf("Unable to retrieve calendars list: %v", err)
	}

	for _, calendarItem := range calendars.Items {
		log.Printf("[%s] %s", calendarItem.Id, calendarItem.Summary)
	}

	return
}
