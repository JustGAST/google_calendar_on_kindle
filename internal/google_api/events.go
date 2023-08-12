package google_api

import (
	"log"
	"time"

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

func GetEventsForMonth(srv *calendar.Service, calendarId string) []*calendar.Event {
	events, err := srv.Events.List(calendarId).
		TimeMin(getStartDay().Format(time.RFC3339)).
		SingleEvents(true).
		OrderBy("startTime").
		TimeMax(getStartDay().AddDate(0, 0, 42).Format(time.RFC3339)).
		Do()

	if err != nil {
		log.Fatalf("Unable to retrieve events: %v", err)
	}

	for _, event := range events.Items {
		log.Printf("[%s %s] %s", event.Start.Date, event.Start.DateTime, event.Summary)
	}

	return events.Items
}

func getStartDay() time.Time {
	t := time.Now().AddDate(0, 0, -time.Now().Day()+1)
	return t.AddDate(0, 0, -int(t.Weekday())+1)
}
