package calendar_image

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

	"github.com/fogleman/gg"

	"github.com/justgast/google_calendar_on_kindle/internal/calendar_utils"
)

const imgWidth = 800
const imgHeight = 600
const headerHeight = 15
const cols = 7
const rows = 6

var colWidth float64 = imgWidth/cols + 1
var rowHeight float64 = ((imgHeight - headerHeight) / rows) + 1

var daysOfWeek = map[int]string{
	0: "Пн",
	1: "Вт",
	2: "Ср",
	3: "Чт",
	4: "Пт",
	5: "Сб",
	6: "Вс",
}

func DrawCalendar(eventsByDay FormattedDayEvents) error {
	var img = gg.NewContext(imgWidth, imgHeight)

	if err := img.LoadFontFace("/Library/Fonts/Arial Unicode.ttf", 10); err != nil {
		return fmt.Errorf("error loading font: %w", err)
	}

	img.SetHexColor("fff")
	img.Clear()

	img.SetHexColor("000")
	img.SetLineWidth(1)

	img.DrawRectangle(0, 0, imgWidth, headerHeight)

	for col := float64(0); col < cols; col++ {
		img.DrawRectangle(colWidth*col, 0, colWidth, imgHeight)
		img.DrawString(daysOfWeek[int(col)], colWidth*col+colWidth/2.2, headerHeight*0.8)
	}

	for row := float64(0); row < rows; row++ {
		img.DrawRectangle(0, rowHeight*row+headerHeight, imgWidth, rowHeight)
	}

	err := drawEvents(eventsByDay, img)
	if err != nil {
		return fmt.Errorf("error drawing events: %w", err)
	}

	img.Stroke()

	var imgBuffer bytes.Buffer
	err = img.EncodePNG(&imgBuffer)
	if err != nil {
		return fmt.Errorf("error encoding image: %w", err)
	}

	manipulatedImg, err := png.Decode(&imgBuffer)
	if err != nil {
		return fmt.Errorf("error decoding image: %w", err)
	}

	var result image.Image = rgbToGray(manipulatedImg)
	result = rotate(result)

	outFile, err := os.Create("calendar.png")
	err = png.Encode(outFile, result)
	if err != nil {
		return fmt.Errorf("error saving image: %w", err)
	}

	return nil
}

func drawEvents(events FormattedDayEvents, img *gg.Context) error {
	currentCol := 0
	currentRow := 0
	cells := cols * rows
	startDay := calendar_utils.GetStartDay()
	currentDay := startDay
	for i := 0; i < cells; i++ {
		if currentCol == cols {
			currentCol = 0
			currentRow++
		}

		day := currentDay.Format("02")
		key := currentDay.Format("0102")

		img.DrawString(day, float64(currentCol)*colWidth+2, headerHeight+float64(currentRow)*rowHeight+12)

		if dayEvents, ok := events[key]; ok {
			eventsString := []string{}
			for _, event := range dayEvents {
				eventsString = append(eventsString, event.Time+" "+event.Summary)
			}

			img.DrawStringWrapped(
				strings.Join(eventsString, "\n"),
				float64(currentCol)*colWidth+2,
				headerHeight+16+float64(currentRow)*rowHeight,
				0.0,
				0.0,
				colWidth,
				1.8,
				gg.AlignLeft,
			)

		}

		currentDay = currentDay.AddDate(0, 0, 1)
		currentCol++
	}

	return nil
}

func rgbToGray(img image.Image) *image.Gray {
	bounds := img.Bounds()
	grayImage := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor)
			grayImage.Set(x, y, grayColor)
		}
	}

	return grayImage
}

func rotate(img image.Image) image.Image {
	bounds := img.Bounds()
	rotatedImage := image.NewGray(image.Rect(0, 0, bounds.Dy(), bounds.Dx()))

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldColor := img.At(x, y)
			rotatedImage.Set(bounds.Max.Y-y-1, x, oldColor)
		}
	}

	return rotatedImage
}
