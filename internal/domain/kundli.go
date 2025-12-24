package domain

type Kundli struct {
	Grahas []Graha
	Bhavas []Bhava
	Dasha  map[string]Dasha
}
