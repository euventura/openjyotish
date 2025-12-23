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

func newCalc(BirthdateTime time.Time, Lat, Lng float64) *Calc {
	return &Calc{
		BirthdateTime: BirthdateTime.UTC(),
		BirthLocation: LatLng{
			Lat: Lat,
			Lng: Lng,
		},
	}
}
