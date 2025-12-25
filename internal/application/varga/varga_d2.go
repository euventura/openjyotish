package varga

import (
	"openjyotish/internal/application"
	"openjyotish/internal/domain"
)

type D2Varga struct {
	domain.IVargaCalc
}

func newD2Varga(k *application.KundliService) domain.IVargaCalc {
	return &D2Varga{}
}

func (d2 *D2Varga) Name() string {
	return "d2"
}

func (d2 *D2Varga) Division() int {
	return 2
}

func (d2 *D2Varga) Calculate(info *domain.VargaCalc) (domain.Varga, error) {

	var varga domain.Varga
	var newGrahas []domain.Graha
	lagna := info.Kundli.Bhavas[0]
	lagnaRashi := 5

	if (lagna.Rashi%2 == 0 && lagna.Degree > 15) || (lagna.Rashi%2 != 0 && lagna.Degree <= 15) {
		lagnaRashi = 4
	}

	for _, graha := range info.Kundli.Grahas {
		rashi := 5

		if (graha.Rashi%2 == 0 && graha.Lng > 15) || (graha.Rashi%2 != 0 && graha.Lng <= 15) {
			rashi = 4
		}

		bhava := 2
		if rashi == lagnaRashi {
			bhava = 1
		}

		newGrahas = append(newGrahas, domain.Graha{
			Name:        graha.Name,
			EName:       graha.EName,
			Lat:         graha.Lat,
			Lng:         graha.Lng,
			Dec:         graha.Dec,
			Speed:       graha.Speed,
			Rashi:       rashi,
			BhavaNumber: bhava,
			Degree:      graha.Degree,
		})

	}
	varga.Grahas = newGrahas
	varga.Lagna = lagna
	varga.Name = d2.Name()
	varga.Division = d2.Division()
	return varga, nil
}
