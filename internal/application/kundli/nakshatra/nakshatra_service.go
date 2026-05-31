package nakshatra

import (
	"fmt"
	"openjyotish/internal/domain"
)

type NakshatraService struct{}

func NewNakshatraService() *NakshatraService {
	return &NakshatraService{}
}

func (ns *NakshatraService) CalcNakshatra(degree float64) (domain.Nakshatra, error) {
	for _, nakshatra := range domain.Nakshatras {
		if degree >= nakshatra.Start && degree < nakshatra.End {
			nakshatra.Pada = int((degree-nakshatra.Start)/3.2) + 1
			return nakshatra, nil
		}
	}
	return domain.Nakshatra{}, fmt.Errorf("invalid degree")
}

func (ns *NakshatraService) FillNakshatra(k *domain.Kundli) {
	nGrahas := []domain.Graha{}
	for _, g := range k.Grahas {
		g.Nakshatra, _ = ns.CalcNakshatra(g.Lng)
		nGrahas = append(nGrahas, g)
	}
	k.Grahas = nGrahas
}
