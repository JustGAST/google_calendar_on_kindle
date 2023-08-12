package main

import (
	"context"
	"github.com/justgast/google_calendar_on_kindle/internal/calendar_image"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"log"
	"os"

	"github.com/justgast/google_calendar_on_kindle/internal/google_api"
)

const calendarId = "u80gdf74rouv41pl2a10r08ji0@group.calendar.google.com"

func main() {
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

	log.Printf("formattedEvents123: %+v \n", formattedEvents)

	return

	//t := time.Now().Truncate(time.Hour * 24 * 28).Format(time.RFC3339)
	//log.Println(t)
	//events, err := srv.Events.List("primary").ShowDeleted(false).
	//	SingleEvents(true).TimeMin(t).OrderBy("startTime").Do()
	//if err != nil {
	//	log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	//}
	//fmt.Println("Upcoming events:")
	//if len(events.Items) == 0 {
	//	fmt.Println("No upcoming events found.")
	//} else {
	//	for _, item := range events.Items {
	//		date := item.Start.DateTime
	//		if date == "" {
	//			date = item.Start.Date
	//		}
	//		fmt.Printf("%v (%v)\n", item.Summary, date)
	//	}
	//}
}
