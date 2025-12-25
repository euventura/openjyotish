package varga

import "openjyotish/internal/domain"

type VargaService struct {
	Calculators map[string]domain.IVargaCalc
}

func NewVargaService() *VargaService {
	vs := &VargaService{
		Calculators: make(map[string]domain.IVargaCalc),
	}
	vs.registerVargas()
	return vs
}

func (vs *VargaService) registerVargas() {

	vs.Calculators["d2"] = &D2Varga{}
	vs.Calculators["d4"] = &D4Varga{}
	// "d3",
	// "d4",
	// "d7",
	// "d9",
	// "d10",
	// "d12",
	// "d16",
	// "d20",
	// "d24",
	// "d27",
	// "d30",
	// "d40",
	// "d45",
	// "d60",

}

func (vs *VargaService) CalculateVargas() []string {
	var names []string
	for name := range vs.Calculators {
		names = append(names, name)
	}
	return names
}
