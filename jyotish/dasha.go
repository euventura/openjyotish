package jyotish

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
	Moon  Graha
	Lagna Graha
}
