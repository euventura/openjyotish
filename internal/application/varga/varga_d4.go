package varga

import (
	"math"
	"openjyotish/internal/application"
	"openjyotish/internal/domain"
)

type D4Varga struct {
	domain.IVargaCalc
}

func newD4Varga(k *application.KundliService) domain.IVargaCalc {
	return &D4Varga{}
}

func (d2 *D4Varga) Name() string {
	return "d4"
}

func (d2 *D4Varga) Division() int {
	return 4
}

func (d2 *D4Varga) Calculate(info *domain.VargaCalc) (domain.Varga, error) {

	var varga domain.Varga
	var newGrahas []domain.Graha
	lagna := info.Kundli.Bhavas[0]

	lagnaRashi := lagna.Rashi

	lagnaPlace := int(math.Floor(lagna.Degree / (30 / 4)))

	switch lagnaPlace {
	case 1:
		lagnaRashi += 4
	case 2:
		lagnaRashi += 7
	case 3:
		lagnaRashi += 10
	}

	if lagnaRashi > 12 {
		lagnaRashi = lagnaRashi - 12
	}

	lagna.Rashi = lagnaRashi

	for _, graha := range info.Kundli.Grahas {
		rashi := graha.Rashi

		place := int(math.Floor(graha.Degree / (30 / 4)))

		switch place {
		case 1:
			rashi += 4
		case 2:
			rashi += 7
		case 3:
			rashi += 10
		}

		if rashi > 12 {
			rashi = rashi - 12
		}

		newGrahas = append(newGrahas, domain.Graha{
			Name:        graha.Name,
			EName:       graha.EName,
			Lat:         graha.Lat,
			Lng:         graha.Lng,
			Dec:         graha.Dec,
			Speed:       graha.Speed,
			Rashi:       rashi,
			BhavaNumber: 1,
			Degree:      graha.Degree,
		})

	}

	varga.Grahas = newGrahas
	varga.Lagna = lagna
	varga.Name = d2.Name()
	varga.Division = d2.Division()
	return varga, nil
}
