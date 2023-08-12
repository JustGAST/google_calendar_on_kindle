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
		startTime, err := getEventTime(event.Start)
		if err != nil {
			fmt.Println("error parsing event start time:", err)
			continue
		}

		endTime, err := getEventTime(event.End)
		if err != nil {
			fmt.Println("error parsing event end time:", err)
			continue
		}

		mapKey := startTime.Format("0102")
		formattedEvents[mapKey] = append(formattedEvents[mapKey], FormattedDayEvent{
			Time:    startTime.Format("15:04"),
			Summary: event.Summary,
		})

		newStartTime := startTime
		for newStartTime.Day() != endTime.Day() {
			newStartTime = newStartTime.AddDate(0, 0, 1)
			mapKey = newStartTime.Format("0102")

			formattedEvents[mapKey] = append(formattedEvents[mapKey], FormattedDayEvent{
				Time:    startTime.Format("15:04"),
				Summary: event.Summary,
			})
		}
	}

	return formattedEvents
}

func getEventTime(eventDateTime *calendar.EventDateTime) (time.Time, error) {
	var eventTime time.Time
	var err error
	if eventDateTime.DateTime != "" {
		eventTime, err = time.Parse(time.RFC3339, eventDateTime.DateTime)
		if err != nil {
			return time.Time{}, fmt.Errorf("unable to parse event time: %w", err)
		}
	}

	if eventDateTime.Date != "" {
		eventTime, err = time.Parse(time.RFC3339, eventDateTime.Date+"T00:00:00+03:00")
		if err != nil {
			return time.Time{}, fmt.Errorf("unable to parse event time: %w", err)
		}
	}

	return eventTime, nil
}
