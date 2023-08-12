package calendar_image

import (
	"fmt"
	"github.com/fogleman/gg"
)

const imgWidth = 800
const imgHeight = 600
const headerHeight = 15

var daysOfWeek = map[int]string{
	0: "Пн",
	1: "Вт",
	2: "Ср",
	3: "Чт",
	4: "Пт",
	5: "Сб",
	6: "Вс",
}

func DrawCalendar() error {
	var img = gg.NewContext(imgWidth, imgHeight)

	if err := img.LoadFontFace("/Library/Fonts/Arial Unicode.ttf", 12); err != nil {
		return fmt.Errorf("error loading font: %w", err)
	}

	img.SetHexColor("fff")
	img.Clear()

	img.SetHexColor("000")
	img.SetLineWidth(1)

	img.DrawRectangle(0, 0, imgWidth, headerHeight)

	for col := float64(0); col < 7; col++ {
		var colWidth float64 = imgWidth/7 + 1
		img.DrawRectangle(colWidth*col, 0, colWidth, imgHeight)
		img.DrawString(daysOfWeek[int(col)], colWidth*col+colWidth/2.2, headerHeight*0.8)
	}

	for row := float64(0); row < 6; row++ {
		var rowHeight float64 = ((imgHeight - headerHeight) / 6) + 1
		img.DrawRectangle(0, rowHeight*row+headerHeight, imgWidth, rowHeight)
	}

	img.Stroke()

	err := img.SavePNG("calendar.png")
	if err != nil {
		return fmt.Errorf("error saving png: %w", err)
	}

	return nil
}
