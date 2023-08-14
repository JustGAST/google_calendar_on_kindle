package use_cases

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"github.com/justgast/google_calendar_on_kindle/internal/calendar_image"
	"github.com/justgast/google_calendar_on_kindle/internal/google_api"
)

func GenerateCalendarImage(calendarId string) {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := google_api.GetClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	events := google_api.GetEventsForMonth(srv, calendarId)
	formattedEvents := calendar_image.FormatEvents(events)

	err = calendar_image.DrawCalendar(formattedEvents)
	if err != nil {
		log.Fatalf("Unable to draw calendar: %v", err)
	}
}
