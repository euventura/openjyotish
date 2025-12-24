package jyotish

import (
	"math"
)

var engToSanskrit map[string]string = map[string]string{
	"Moon":    "Chandra",
	"Mars":    "Mangala",
	"Mercury": "Budha",
	"Jupiter": "Guru",
	"Venus":   "Shukra",
	"Saturn":  "Shani",
	"Rahu":    "Rahu",
	"Ketu":    "Ketu",
}

func grahaBuild(graha *Graha, lagna Bhava) {

	// lagRas := int(math.Floor((lagna.RawDegree / 30))) + 1
	graha.BhavaNumber = int(math.Floor(math.Mod((graha.Lng-lagna.RawDegree), 360)/30)) + 1

	if lagna.RawDegree > graha.Lng {
		graha.BhavaNumber = 13 - int(math.Floor(math.Mod((lagna.RawDegree-graha.Lng), 360)/30)+1)
	}

	graha.Rashi = int(math.Floor((graha.Lng / 30))) + 1
	graha.Name = engToSanskrit[graha.Name]
}
