package application

import (
	"fmt"
	"openjyotish/internal/domain"
)

type NakshatraService struct{}

func (ns *NakshatraService) CalcNakshatra(degree float64, nakshatras []domain.Nakshatra) (domain.Nakshatra, error) {
	for _, nakshatra := range nakshatras {
		if degree >= nakshatra.Start && degree < nakshatra.End {
			nakshatra.Pada = int((degree-nakshatra.Start)/3.2) + 1
			return nakshatra, nil
		}
	}
	return domain.Nakshatra{}, fmt.Errorf("invalid degree")
}

func (ns *NakshatraService) FillNakshatra(k *domain.Kundli, nakshatras []domain.Nakshatra) {
	nGrahas := []domain.Graha{}
	for _, g := range k.Grahas {
		g.Nakshatra, _ = ns.CalcNakshatra(g.Lng, nakshatras)
		nGrahas = append(nGrahas, g)
	}
	k.Grahas = nGrahas
}
