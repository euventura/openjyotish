package dasha

import (
	"fmt"
	"openjyotish/internal/domain"
	"time"
)

type DashaService struct {
	calculators map[string]domain.DashaCalculator
}

func NewDashaService() *DashaService {
	service := &DashaService{
		calculators: make(map[string]domain.DashaCalculator),
	}
	service.registerCalculators()
	return service
}

func (ds *DashaService) registerCalculators() {
	ds.calculators["Vimsottari"] = &VimsottariCalculator{}
}

func (ds *DashaService) RegisterCalculator(calculator domain.DashaCalculator) {
	ds.calculators[calculator.Name()] = calculator
}

func (ds *DashaService) CalculateDasha(k *domain.Kundli, birthDate time.Time, dashaName string) error {

	if k.Dashas == nil {
		k.Dashas = make(map[string]domain.Dasha)
	}

	var moon, lagna domain.Graha
	for _, g := range k.Grahas {
		if g.Name == "Chandra" {
			moon = g
		}
	}

	calculator, ok := ds.calculators[dashaName]
	if !ok {
		return fmt.Errorf("dasha calculator for %s not found", dashaName)
	}

	dashaCalcInfo := domain.DashaCalc{
		Moon:      moon,
		Lagna:     lagna,
		BirthDate: birthDate,
	}

	dasha, err := calculator.Calculate(dashaCalcInfo)
	if err != nil {
		return err
	}

	k.Dashas[dashaName] = dasha

	return nil
}

func (ds *DashaService) CalculateAllDashas(k *domain.Kundli, birthDate time.Time) error {
	for dashaName := range ds.calculators {
		if err := ds.CalculateDasha(k, birthDate, dashaName); err != nil {
			return fmt.Errorf("failed to calculate %s: %w", dashaName, err)
		}
	}
	return nil
}
