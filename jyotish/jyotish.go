package jyotish

import (
	"math"
	"openjyotish/swiss"
)

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
}

type Bhava struct {
	Number    int
	Name      string
	Rashi     int
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

	for i, sBhav := range swiss.Bhavas {

		bRashi := 1

		if i > 0 && swiss.Bhavas[0].Start > sBhav.Start {
			bRashi = 13 - int(math.Floor(math.Mod((swiss.Bhavas[0].Start-sBhav.Start), 360)/30)+1)
		}

		if swiss.Bhavas[0].Start < sBhav.Start && i > 0 {
			bRashi = int(math.Floor(math.Mod((sBhav.Start-swiss.Bhavas[0].Start), 360)/30)) + 1

		}

		bhavas = append(bhavas, Bhava{Number: sBhav.Number, Start: sBhav.Start, End: sBhav.End, Rashi: bRashi})
	}

	for _, sGrah := range swiss.Grahas {

		newGraha := Graha{
			EName: sGrah.Name,
			Lat:   sGrah.Latitude,
			Lng:   sGrah.Longitude,
			Dec:   sGrah.Dec,
			Speed: sGrah.Speed,
		}

		grahaBuild(&newGraha, bhavas[0])
		grahas = append(grahas, newGraha)

	}

}
