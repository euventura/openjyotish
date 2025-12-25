package domain

type Rashi struct {
	Name   string
	EName  string
	Number int
	Grahas []Graha
}

var Rashis = []Rashi{
	{Name: "Mesha", EName: "Aries", Number: 1},
	{Name: "Vrishabha", EName: "Taurus", Number: 2},
	{Name: "Mithuna", EName: "Gemini", Number: 3},
	{Name: "Karka", EName: "Cancer", Number: 4},
	{Name: "Simha", EName: "Leo", Number: 5},
	{Name: "Kanya", EName: "Virgo", Number: 6},
	{Name: "Tula", EName: "Libra", Number: 7},
	{Name: "Vrischika", EName: "Scorpio", Number: 8},
	{Name: "Dhanu", EName: "Sagittarius", Number: 9},
	{Name: "Makara", EName: "Capricorn", Number: 10},
	{Name: "Kumbha", EName: "Aquarius", Number: 11},
	{Name: "Meena", EName: "Pisces", Number: 12},
}
