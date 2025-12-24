package swiss

import "time"

type Calc struct {
	BirthdateTime time.Time
	BirthLocation LatLng
	MeanNode      bool
	Ayananmsa     int
	BhavaSystem   string
}

type LatLng struct {
	Lat float64
	Lng float64
}

func execute(birthdateTime time.Time, lat, lng float64, BhavaSys string, ayanamsa string, trueNode bool) (Result, error) {

	return ExecSwiss(&SwissOptions{
		DateTime:    birthdateTime,
		BhavaSystem: BhavaSys,
		Geopos:      LatLng{Lat: lat, Lng: lng},
		Ayanamsa:    ayanamsa,
		TrueNode:    trueNode,
	})
}
