package dasha

import (
	"fmt"
	"math"
	"openjyotish/internal/application/nakshatra"
	"openjyotish/internal/domain"
	"strings"
	"time"
)

type VimsottariCalculator struct{}

func (vc *VimsottariCalculator) Name() string {
	return "Vimsottari"
}

func (vc *VimsottariCalculator) Calculate(info domain.DashaCalc) (domain.Dasha, error) {
	moonLng := info.Moon.Lng
	nakshatraSvc := nakshatra.NewNakshatraService()
	nakshatraData, err := nakshatraSvc.CalcNakshatra(moonLng)
	if err != nil {
		return domain.Dasha{}, fmt.Errorf("could not calculate nakshatra: %w", err)
	}

	dashaLord := nakshatraData.Isha
	dashaDuration := float64(vimsottariDashaSizes[dashaLord])
	nakshatraSpan := nakshatraData.End - nakshatraData.Start
	moonTraversed := moonLng - nakshatraData.Start
	remainingProportion := (nakshatraSpan - moonTraversed) / nakshatraSpan
	remainingYears := remainingProportion * dashaDuration

	totalDaysForDasha := remainingYears * 365.25
	endDate := info.BirthDate.AddDate(0, 0, int(totalDaysForDasha))

	periods := []domain.DashaPeriod{}
	periods = append(periods, domain.DashaPeriod{
		Graha: domain.Graha{Name: dashaLord},
		Start: info.BirthDate,
		End:   endDate,
		Antar: vc.calculateAntardashas(dashaLord, info.BirthDate, endDate, 1),
	})

	currentDate := endDate
	lordIndex := vc.findLordIndex(dashaLord)
	if lordIndex == -1 {
		return domain.Dasha{}, fmt.Errorf("invalid dasha lord: %s", dashaLord)
	}

	for i := 1; i < len(vimsottariLordSequence); i++ {
		nextLordIndex := (lordIndex + i) % len(vimsottariLordSequence)
		nextLord := vimsottariLordSequence[nextLordIndex]
		duration := vimsottariDashaSizes[nextLord]

		startDate := currentDate
		totalDays := int(math.Round(float64(duration) * 365.25))
		endDate := startDate.AddDate(0, 0, totalDays)

		periods = append(periods, domain.DashaPeriod{
			Graha: domain.Graha{Name: nextLord},
			Start: startDate,
			End:   endDate,
			Antar: vc.calculateAntardashas(nextLord, startDate, endDate, 1),
		})
		currentDate = endDate
	}

	return domain.Dasha{Name: "Vimsottari", Periods: periods}, nil
}

func (vc *VimsottariCalculator) calculateAntardashas(mainLord string, startDate time.Time, endDate time.Time, level int) []domain.DashaPeriod {
	if level > 1 {
		return []domain.DashaPeriod{}
	}

	antarPeriods := []domain.DashaPeriod{}
	mainLordIndex := vc.findLordIndex(mainLord)
	if mainLordIndex == -1 {
		return antarPeriods
	}

	totalDays := endDate.Sub(startDate).Hours() / 24
	mainLordDuration := float64(vimsottariDashaSizes[mainLord]) * 365.25

	currentDate := startDate
	for i := 0; i < len(vimsottariLordSequence); i++ {
		antarLordIndex := (mainLordIndex + i) % len(vimsottariLordSequence)
		antarLord := vimsottariLordSequence[antarLordIndex]
		antarLordDuration := float64(vimsottariDashaSizes[antarLord])

		proportionalDays := (antarLordDuration / mainLordDuration) * totalDays
		antarEndDate := currentDate.AddDate(0, 0, int(math.Round(proportionalDays)))

		if currentDate.After(endDate) {
			break
		}

		if antarEndDate.After(endDate) {
			antarEndDate = endDate
		}

		isLastAntardasha := i == len(vimsottariLordSequence)-1
		if isLastAntardasha {
			antarEndDate = endDate
		}

		antarPeriods = append(antarPeriods, domain.DashaPeriod{
			Graha: domain.Graha{Name: antarLord},
			Start: currentDate,
			End:   antarEndDate,
			Antar: []domain.DashaPeriod{},
		})

		currentDate = antarEndDate
	}

	return antarPeriods
}

func (vc *VimsottariCalculator) findLordIndex(lord string) int {
	for i, l := range vimsottariLordSequence {
		if strings.EqualFold(l, lord) {
			return i
		}
	}
	return -1
}

var vimsottariLordSequence = []string{
	"Ketu", "Shukra", "Surya", "Chandra", "Mangala", "Rahu", "Guru", "Shani", "Budha",
}

var vimsottariDashaSizes = map[string]int{
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
