package calendar_image

import (
	"fmt"
	"google.golang.org/api/calendar/v3"
	"time"
)

type FormattedDayEvent struct {
	Time    string
	Summary string
}

type FormattedDayEvents map[string][]FormattedDayEvent

func FormatEvents(events []*calendar.Event) FormattedDayEvents {
	formattedEvents := FormattedDayEvents{}
	for _, event := range events {
		eventTime, err := getEventTime(event)
		if err != nil {
			fmt.Println(err)
			continue
		}

		mapKey := eventTime.Format("0102")
		formattedEvents[mapKey] = append(formattedEvents[mapKey], FormattedDayEvent{
			Time:    eventTime.Format("15:04"),
			Summary: event.Summary,
		})
	}

	return formattedEvents
}

func getEventTime(event *calendar.Event) (time.Time, error) {
	var eventTime time.Time
	var err error
	if event.Start.DateTime != "" {
		eventTime, err = time.Parse(time.RFC3339, event.Start.DateTime)
		if err != nil {
			return time.Time{}, fmt.Errorf("unable to parse event time: %w", err)
		}
	}

	if event.Start.Date != "" {
		eventTime, err = time.Parse(time.RFC3339, event.Start.Date+"T00:00:00+03:00")
		if err != nil {
			return time.Time{}, fmt.Errorf("unable to parse event time: %w", err)
		}
	}

	return eventTime, nil
}
