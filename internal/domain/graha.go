package domain

type Graha struct {
	Name        string
	EName       string
	Rashi       int
	BhavaNumber int
	Bhava       Bhava
	Lat         float64
	Lng         float64
	Dec         float64
	Speed       float64
	Nakshatra   Nakshatra
	Degree      float64
}
