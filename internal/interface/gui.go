package interfaces

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"openjyotish/internal/application"
	"openjyotish/internal/infrastructure"
	"openjyotish/swiss"
)

func RunGUI() {
	myApp := app.New()
	myWindow := myApp.NewWindow("OpenJyotish - Vedic Astrology Calculator")
	myWindow.Resize(fyne.NewSize(800, 600))

	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("DD/MM/YYYY")
	dateEntry.SetText("13/06/1988")

	timeEntry := widget.NewEntry()
	timeEntry.SetPlaceHolder("HH:MM")
	timeEntry.SetText("09:30")

	latEntry := widget.NewEntry()
	latEntry.SetPlaceHolder("Decimal latitude")
	latEntry.SetText("-23.5505")

	lngEntry := widget.NewEntry()
	lngEntry.SetPlaceHolder("Decimal longitude")
	lngEntry.SetText("-46.6333")

	ayanamsaEntry := widget.NewEntry()
	ayanamsaEntry.SetPlaceHolder("Ayanamsa code")
	ayanamsaEntry.SetText("1")

	bhavaEntry := widget.NewEntry()
	bhavaEntry.SetPlaceHolder("Bhava system code")
	bhavaEntry.SetText("P")

	trueNodeCheck := widget.NewCheck("Use True Node", nil)

	resultText := widget.NewMultiLineEntry()
	resultText.Disable()
	resultText.SetPlaceHolder("Results will appear here...")

	calculateBtn := widget.NewButton("Calculate", func() {
		resultText.SetText("Calculating...")

		dateTime, err := time.Parse("02/01/2006 15:04", dateEntry.Text+" "+timeEntry.Text)
		if err != nil {
			resultText.SetText(fmt.Sprintf("Error: invalid date/time format: %v", err))
			return
		}

		lat := 0.0
		lng := 0.0
		fmt.Sscanf(latEntry.Text, "%f", &lat)
		fmt.Sscanf(lngEntry.Text, "%f", &lng)

		opt := &swiss.SwissOptions{
			DateTime:    dateTime,
			BhavaSystem: bhavaEntry.Text,
			Geopos:      swiss.LatLng{Lat: lat, Lng: lng},
			Ayanamsa:    ayanamsaEntry.Text,
			TrueNode:    trueNodeCheck.Checked,
		}

		adapter := infrastructure.SwissAdapter{}
		res, err := adapter.Exec(opt)
		if err != nil {
			resultText.SetText(fmt.Sprintf("Error: swiss calculation failed: %v", err))
			return
		}

		ks := application.NewKundliService(&application.NakshatraService{}, &application.DashaService{})
		k := ks.LoadFromSwiss(res, dateTime)

		output := "Vimsottari Dasha:\n\n"
		if d, ok := k.Dasha["Vimsottari"]; ok {
			for _, p := range d.Periods {
				output += fmt.Sprintf("- %s: %s to %s\n",
					p.Graha.Name,
					p.Start.Format("2006-01-02"),
					p.End.Format("2006-01-02"))
			}
		} else {
			output = "Vimsottari Dasha not computed"
		}

		resultText.SetText(output)
	})

	form := container.NewVBox(
		widget.NewLabel("Birth Date:"),
		dateEntry,
		widget.NewLabel("Birth Time:"),
		timeEntry,
		widget.NewLabel("Latitude:"),
		latEntry,
		widget.NewLabel("Longitude:"),
		lngEntry,
		widget.NewLabel("Ayanamsa:"),
		ayanamsaEntry,
		widget.NewLabel("Bhava System:"),
		bhavaEntry,
		trueNodeCheck,
		calculateBtn,
		widget.NewSeparator(),
		widget.NewLabel("Results:"),
		resultText,
	)

	myWindow.SetContent(container.NewScroll(form))
	myWindow.ShowAndRun()
}
