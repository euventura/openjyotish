package application

import (
	"math"

	"openjyotish/internal/application/dasha"
	"openjyotish/internal/application/nakshatra"
	"openjyotish/internal/domain"
	"openjyotish/swiss"
	"time"
)

type KundliService struct {
	nakshatraSvc *nakshatra.NakshatraService
	dashaSvc     *dasha.DashaService
}

func NewKundliService() *KundliService {
	ns := nakshatra.NewNakshatraService()
	ds := dasha.NewDashaService()
	return &KundliService{nakshatraSvc: ns, dashaSvc: ds}
}

func (ks *KundliService) LoadFromSwiss(swissRes *swiss.Result, dateTime time.Time) *domain.Kundli {
	k := domain.Kundli{}
	bhavas := []domain.Bhava{}
	for i, sBhav := range swissRes.Bhavas {
		bRashi := 1
		if i > 0 && swissRes.Bhavas[0].Start > sBhav.Start {
			bRashi = 13 - int(math.Floor(math.Mod((swissRes.Bhavas[0].Start-sBhav.Start), 360)/30)+1)
		}
		if swissRes.Bhavas[0].Start < sBhav.Start && i > 0 {
			bRashi = int(math.Floor(math.Mod((sBhav.Start-swissRes.Bhavas[0].Start), 360)/30)) + 1
		}
		bhavas = append(bhavas, domain.Bhava{
			Number:    sBhav.Number,
			Start:     sBhav.Start,
			End:       sBhav.End,
			Rashi:     bRashi,
			RawDegree: sBhav.Start,
			Degree:    math.Mod(sBhav.Start-(float64(sBhav.Number-1)*30), 360),
		})
	}
	k.Bhavas = bhavas

	grahas := []domain.Graha{}
	for _, sGrah := range swissRes.Grahas {
		newGraha := domain.Graha{
			EName: sGrah.Name,
			Lat:   sGrah.Latitude,
			Lng:   sGrah.Longitude,
			Dec:   sGrah.Dec,
			Speed: sGrah.Speed,
		}
		ks.grahaBuild(&newGraha, bhavas[0])
		grahas = append(grahas, newGraha)
	}
	k.Grahas = grahas
	ks.processKundli(&k, dateTime)
	return &k
}

func (ks *KundliService) processKundli(k *domain.Kundli, dateTime time.Time) {
	if ks.nakshatraSvc != nil {
		ks.nakshatraSvc.FillNakshatra(k)
	}
	if ks.dashaSvc != nil {
		_ = ks.dashaSvc.CalculateAllDashas(k, dateTime)
	}
}

func (ks *KundliService) grahaBuild(graha *domain.Graha, lagna domain.Bhava) {
	graha.BhavaNumber = int(math.Floor(math.Mod((graha.Lng-lagna.RawDegree), 360)/30)) + 1
	if lagna.RawDegree > graha.Lng {
		graha.BhavaNumber = 13 - int(math.Floor(math.Mod((lagna.RawDegree-graha.Lng), 360)/30)+1)
	}
	graha.Rashi = int(math.Floor((graha.Lng / 30))) + 1
	if name, ok := domain.EngToSanskrit[graha.EName]; ok {
		graha.Name = name
	} else {
		graha.Name = graha.EName
	}
}
