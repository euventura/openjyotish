package application

import (
	"fmt"
	"math"
	"openjyotish/internal/domain"
	"strings"
	"time"
)

type DashaService struct{}

func (ds *DashaService) CalculateDashas(k *domain.Kundli, birthDate time.Time) error {
	var moon, lagna domain.Graha
	for _, g := range k.Grahas {
		if g.EName == "Moon" {
			moon = g
		}
		if g.EName == "Lagna" {
			lagna = g
		}
	}

	dashaCalcInfo := domain.DashaCalc{
		Moon:      moon,
		Lagna:     lagna,
		BirthDate: birthDate,
	}

	vimsottariDasha, err := Vimsottari(dashaCalcInfo)
	if err != nil {
		return err
	}

	if k.Dasha == nil {
		k.Dasha = make(map[string]domain.Dasha)
	}
	k.Dasha[vimsottariDasha.Name] = vimsottariDasha

	return nil
}

var dashaLordSequence = []string{
	"Ketu", "Shukra", "Surya", "Chandra", "Mangala", "Rahu", "Guru", "Shani", "Budha",
}

var dashaSize = map[string]int{
	"Ketu":    7,
	"Shukra":  20,
	"Surya":   6,
	"Chandra": 10,
	"Mangala": 7,
	"Rahu":    18,
	"Guru":    16,
	"Shani":   19,
	"Budha":   17,
}

func findLordIndex(lord string) int {
	for i, l := range dashaLordSequence {
		if strings.EqualFold(l, lord) {
			return i
		}
	}
	return -1
}

func Vimsottari(info domain.DashaCalc) (domain.Dasha, error) {
	moonLng := info.Moon.Lng
	nakshatraSvc := NakshatraService{}
	nakshatra, err := nakshatraSvc.CalcNakshatra(moonLng)
	if err != nil {
		return domain.Dasha{}, fmt.Errorf("could not calculate nakshatra: %w", err)
	}

	dashaLord := nakshatra.Isha
	dashaDuration := float64(dashaSize[dashaLord])
	nakshatraSpan := nakshatra.End - nakshatra.Start
	moonTraversed := moonLng - nakshatra.Start
	remainingProportion := (nakshatraSpan - moonTraversed) / nakshatraSpan
	remainingYears := remainingProportion * dashaDuration

	totalDaysForDasha := remainingYears * 365.25
	endDate := info.BirthDate.AddDate(0, 0, int(totalDaysForDasha))

	periods := []domain.DashaPeriod{}
	periods = append(periods, domain.DashaPeriod{
		Graha: domain.Graha{Name: dashaLord},
		Start: info.BirthDate,
		End:   endDate,
	})

	currentDate := endDate
	lordIndex := findLordIndex(dashaLord)
	if lordIndex == -1 {
		return domain.Dasha{}, fmt.Errorf("invalid dasha lord: %s", dashaLord)
	}

	for i := 1; i < len(dashaLordSequence); i++ {
		nextLordIndex := (lordIndex + i) % len(dashaLordSequence)
		nextLord := dashaLordSequence[nextLordIndex]
		duration := dashaSize[nextLord]

		startDate := currentDate
		totalDays := int(math.Round(float64(duration) * 365.25))
		endDate := startDate.AddDate(0, 0, totalDays)

		periods = append(periods, domain.DashaPeriod{
			Graha: domain.Graha{Name: nextLord},
			Start: startDate,
			End:   endDate,
		})
		currentDate = endDate
	}

	return domain.Dasha{Name: "Vimsottari", Periods: periods}, nil
}
