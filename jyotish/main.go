package jyotish

import "openjyotish/swiss"

type Graha struct {
	Name  string
	EName string
	Rashi string
	Bhava Bhava
	Lat   float64
	Lng   float64
	Dec   float64
	Speed float64
}

type Bhava struct {
	Number    int
	Name      string
	Rashi     string
	Start     float64
	End       float64
	RawDegree float64
}

type Rashi struct {
	Name   string
	EName  string
	Rashi  string
	Grahas []Graha
}

var grahas []Graha
var bhavas []Bhava

func loadFromSwiss(swiss *swiss.Result) {

	for _, sBhav := range swiss.Bhavas {
		bhavas = append(bhavas, Bhava{Number: sBhav.Number, Start: sBhav.Start, End: sBhav.End})
	}

	for _, sGrah := range swiss.Grahas {
		grahas = append(grahas, Graha{
			EName: sGrah.Name,
			Lat:   sGrah.Latitude,
			Lng:   sGrah.Longitude,
			Dec:   sGrah.Dec,
			Speed: sGrah.Speed,
		})
	}

}
