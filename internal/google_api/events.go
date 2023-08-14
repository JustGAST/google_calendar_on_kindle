package google_api

import (
	"github.com/justgast/google_calendar_on_kindle/internal/calendar_utils"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

func GetEventsForMonth(srv *calendar.Service, calendarId string) []*calendar.Event {
	events, err := srv.Events.List(calendarId).
		TimeMin(calendar_utils.GetStartDay().Format(time.RFC3339)).
		SingleEvents(true).
		OrderBy("startTime").
		TimeMax(calendar_utils.GetStartDay().AddDate(0, 0, 42).Format(time.RFC3339)).
		Do()

	if err != nil {
		log.Fatalf("Unable to retrieve events: %v", err)
	}

	return events.Items
}
