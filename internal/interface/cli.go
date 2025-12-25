package interfaces

import (
	"flag"
	"fmt"
	"log"
	"time"

	"openjyotish/internal/application"
	"openjyotish/internal/infrastructure"
	"openjyotish/swiss"
)

func Run() {
	gui := flag.Bool("gui", false, "Run GUI mode")
	dateStr := flag.String("date", "13/06/1988", "Date in DD/MM/YYYY format")
	timeStr := flag.String("time", "09:30", "Time in HH:MM format")
	lat := flag.Float64("lat", -23.5505, "Decimal latitude")
	lng := flag.Float64("lng", -46.6333, "Decimal longitude")
	ayanamsa := flag.String("ayanamsa", "1", "Ayanamsa code")
	bhava := flag.String("bhava", "P", "Bhava system code")
	trueNode := flag.Bool("trueNode", false, "Use true node")
	flag.Parse()

	if *gui {
		RunGUI()
		return
	}

	dateTime, err := time.Parse("02/01/2006 15:04", *dateStr+" "+*timeStr)
	if err != nil {
		log.Fatalf("invalid date/time format: %v", err)
	}

	opt := &swiss.SwissOptions{
		DateTime:    dateTime,
		BhavaSystem: *bhava,
		Geopos:      swiss.LatLng{Lat: *lat, Lng: *lng},
		Ayanamsa:    *ayanamsa,
		TrueNode:    *trueNode,
	}

	adapter := infrastructure.SwissAdapter{}
	res, err := adapter.Exec(opt)
	if err != nil {
		log.Fatalf("swiss calculation failed: %v", err)
	}

	ks := application.NewKundliService(&application.NakshatraService{}, &application.DashaService{})
	k := ks.LoadFromSwiss(res, dateTime)

	if d, ok := k.Dasha["Vimsottari"]; ok {
		fmt.Println("Vimsottari Dasha:")
		for _, p := range d.Periods {
			fmt.Printf("- %s: %s to %s\n", p.Graha.Name, p.Start.Format("2006-01-02"), p.End.Format("2006-01-02"))
		}
	} else {
		fmt.Println("Vimsottari Dasha not computed")
	}
}
