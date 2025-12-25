package domain

type Kundli struct {
	Grahas []Graha
	Bhavas []Bhava
	Dashas map[string]Dasha
}
