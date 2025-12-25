package domain

import "time"

type DashaPeriod struct {
	Title       string
	Graha       Graha
	BhavaNumber int
	Start       time.Time
	End         time.Time
	Antar       []DashaPeriod
}

type Dasha struct {
	Name    string
	Periods []DashaPeriod
}

type DashaCalc struct {
	Moon      Graha
	Lagna     Graha
	BirthDate time.Time
	Name      string
}

type DashaCalculator interface {
	Name() string
	Calculate(info DashaCalc) (Dasha, error)
}
