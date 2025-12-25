package test

import (
	"testing"

	"openjyotish/internal/application/varga"
	"openjyotish/internal/domain"
)

func TestNewVargaService(t *testing.T) {
	vs := varga.NewVargaService()

	if vs == nil {
		t.Fatal("VargaService should not be nil")
	}
}

func TestVargaServiceRegisterD2(t *testing.T) {
	vs := varga.NewVargaService()

	if len(vs.Calculators) == 0 {
		t.Fatal("Calculators should have registered vargas")
	}
}

func TestD2VargaName(t *testing.T) {
	d2 := &varga.D2Varga{}

	expected := "d2"
	if d2.Name() != expected {
		t.Errorf("expected name %q, got %q", expected, d2.Name())
	}
}

func TestD2VargaDivision(t *testing.T) {
	d2 := &varga.D2Varga{}

	expected := 2
	if d2.Division() != expected {
		t.Errorf("expected division %d, got %d", expected, d2.Division())
	}
}

func TestD2VargaCalculate(t *testing.T) {
	tests := []struct {
		name      string
		kundli    *domain.Kundli
		wantError bool
		validate  func(*testing.T, domain.Varga)
	}{
		{
			name: "even rashi lagna with degree > 15",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     2,
						Degree:    20,
						RawDegree: 50,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:  "Sun",
						EName: "Su",
						Rashi: 2,
						Lng:   25,
						Lat:   0,
						Dec:   0,
						Speed: 1,
					},
				},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if v.Name != "d2" {
					t.Errorf("expected name d2, got %s", v.Name)
				}
				if v.Division != 2 {
					t.Errorf("expected division 2, got %d", v.Division)
				}
				if len(v.Grahas) != 1 {
					t.Errorf("expected 1 graha, got %d", len(v.Grahas))
				}
				if v.Grahas[0].Rashi != 4 {
					t.Errorf("expected rashi 4, got %d", v.Grahas[0].Rashi)
				}
			},
		},
		{
			name: "odd rashi lagna with degree <= 15",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     1,
						Degree:    10,
						RawDegree: 10,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:  "Sun",
						EName: "Su",
						Rashi: 1,
						Lng:   10,
						Lat:   0,
						Dec:   0,
						Speed: 1,
					},
				},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if v.Grahas[0].Rashi != 4 {
					t.Errorf("expected rashi 4, got %d", v.Grahas[0].Rashi)
				}
			},
		},
		{
			name: "multiple grahas",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     2,
						Degree:    20,
						RawDegree: 50,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:  "Sun",
						EName: "Su",
						Rashi: 2,
						Lng:   25,
						Lat:   0,
						Dec:   0,
						Speed: 1,
					},
					{
						Name:  "Moon",
						EName: "Mo",
						Rashi: 2,
						Lng:   25,
						Lat:   0,
						Dec:   0,
						Speed: 1,
					},
					{
						Name:  "Mars",
						EName: "Ma",
						Rashi: 3,
						Lng:   15,
						Lat:   0,
						Dec:   0,
						Speed: 1,
					},
				},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if len(v.Grahas) != 3 {
					t.Errorf("expected 3 grahas, got %d", len(v.Grahas))
				}
			},
		},
		{
			name: "lagna and graha in same rashi",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     2,
						Degree:    20,
						RawDegree: 50,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:  "Sun",
						EName: "Su",
						Rashi: 2,
						Lng:   25,
						Lat:   0,
						Dec:   0,
						Speed: 1,
					},
				},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if v.Grahas[0].BhavaNumber != 1 {
					t.Errorf("expected bhava number 1, got %d", v.Grahas[0].BhavaNumber)
				}
			},
		},
		{
			name: "lagna and graha in different rashi",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     2,
						Degree:    20,
						RawDegree: 50,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:  "Sun",
						EName: "Su",
						Rashi: 1,
						Lng:   10,
						Lat:   0,
						Dec:   0,
						Speed: 1,
					},
				},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if v.Grahas[0].BhavaNumber != 1 {
					t.Errorf("expected bhava number 1, got %d", v.Grahas[0].BhavaNumber)
				}
			},
		},
		{
			name: "empty grahas",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     1,
						Degree:    10,
						RawDegree: 10,
					},
				},
				Grahas: []domain.Graha{},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if len(v.Grahas) != 0 {
					t.Errorf("expected 0 grahas, got %d", len(v.Grahas))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d2 := &varga.D2Varga{}
			vargaCalc := &domain.VargaCalc{Kundli: *tt.kundli}

			vargaResult, err := d2.Calculate(vargaCalc)

			if (err != nil) != tt.wantError {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.wantError {
				tt.validate(t, vargaResult)
			}
		})
	}
}

func TestD2VargaCalculateEvenRashi(t *testing.T) {
	tests := []struct {
		name              string
		lagnaRashi        int
		lagnaDegree       float64
		grahaRashi        int
		grahaDegree       float64
		expectedGrahaRash int
	}{
		{name: "even rashi > 15", lagnaRashi: 2, lagnaDegree: 20, grahaRashi: 2, grahaDegree: 20, expectedGrahaRash: 4},
		{name: "even rashi <= 15", lagnaRashi: 2, lagnaDegree: 10, grahaRashi: 2, grahaDegree: 10, expectedGrahaRash: 5},
		{name: "odd rashi > 15", lagnaRashi: 1, lagnaDegree: 20, grahaRashi: 1, grahaDegree: 20, expectedGrahaRash: 5},
		{name: "odd rashi <= 15", lagnaRashi: 1, lagnaDegree: 10, grahaRashi: 1, grahaDegree: 10, expectedGrahaRash: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kundli := &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     tt.lagnaRashi,
						Degree:    tt.lagnaDegree,
						RawDegree: tt.lagnaDegree,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:   "Test",
						EName:  "T",
						Rashi:  tt.grahaRashi,
						Lng:    tt.grahaDegree,
						Lat:    0,
						Dec:    0,
						Speed:  1,
						Degree: tt.grahaDegree,
					},
				},
			}

			d2 := &varga.D2Varga{}
			vargaCalc := &domain.VargaCalc{Kundli: *kundli}
			vargaResult, _ := d2.Calculate(vargaCalc)

			if vargaResult.Grahas[0].Rashi != tt.expectedGrahaRash {
				t.Errorf("expected rashi %d, got %d", tt.expectedGrahaRash, vargaResult.Grahas[0].Rashi)
			}
		})
	}
}

func TestD4VargaName(t *testing.T) {
	d4 := &varga.D4Varga{}

	expected := "d4"
	if d4.Name() != expected {
		t.Errorf("expected name %q, got %q", expected, d4.Name())
	}
}

func TestD4VargaDivision(t *testing.T) {
	d4 := &varga.D4Varga{}

	expected := 4
	if d4.Division() != expected {
		t.Errorf("expected division %d, got %d", expected, d4.Division())
	}
}

func TestD4VargaCalculate(t *testing.T) {
	tests := []struct {
		name      string
		kundli    *domain.Kundli
		wantError bool
		validate  func(*testing.T, domain.Varga)
	}{
		{
			name: "lagna in first place (0-7.5 degrees)",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     1,
						Degree:    5,
						RawDegree: 5,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:   "Sun",
						EName:  "Su",
						Rashi:  1,
						Lng:    5,
						Lat:    0,
						Dec:    0,
						Speed:  1,
						Degree: 5,
					},
				},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if v.Name != "d4" {
					t.Errorf("expected name d4, got %s", v.Name)
				}
				if v.Division != 4 {
					t.Errorf("expected division 4, got %d", v.Division)
				}
				if v.Lagna.Rashi != 1 {
					t.Errorf("expected lagna rashi 1, got %d", v.Lagna.Rashi)
				}
				if v.Grahas[0].Rashi != 1 {
					t.Errorf("expected graha rashi 1, got %d", v.Grahas[0].Rashi)
				}
			},
		},
		{
			name: "lagna in second place (7.5-15 degrees)",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     1,
						Degree:    10,
						RawDegree: 10,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:   "Sun",
						EName:  "Su",
						Rashi:  1,
						Lng:    10,
						Lat:    0,
						Dec:    0,
						Speed:  1,
						Degree: 10,
					},
				},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if v.Lagna.Rashi != 5 {
					t.Errorf("expected lagna rashi 5, got %d", v.Lagna.Rashi)
				}
			},
		},
		{
			name: "lagna in third place (15-22.5 degrees)",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     1,
						Degree:    20,
						RawDegree: 20,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:   "Sun",
						EName:  "Su",
						Rashi:  1,
						Lng:    20,
						Lat:    0,
						Dec:    0,
						Speed:  1,
						Degree: 20,
					},
				},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if v.Lagna.Rashi != 8 {
					t.Errorf("expected lagna rashi 8, got %d", v.Lagna.Rashi)
				}
			},
		},
		{
			name: "lagna in fourth place (22.5-30 degrees)",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     1,
						Degree:    25,
						RawDegree: 25,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:   "Sun",
						EName:  "Su",
						Rashi:  1,
						Lng:    25,
						Lat:    0,
						Dec:    0,
						Speed:  1,
						Degree: 25,
					},
				},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if v.Lagna.Rashi != 11 {
					t.Errorf("expected lagna rashi 11, got %d", v.Lagna.Rashi)
				}
			},
		},
		{
			name: "rashi wrapping when > 12",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     10,
						Degree:    20,
						RawDegree: 20,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:   "Sun",
						EName:  "Su",
						Rashi:  10,
						Lng:    20,
						Lat:    0,
						Dec:    0,
						Speed:  1,
						Degree: 20,
					},
				},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if v.Lagna.Rashi < 1 || v.Lagna.Rashi > 12 {
					t.Errorf("expected rashi between 1-12, got %d", v.Lagna.Rashi)
				}
			},
		},
		{
			name: "multiple grahas with different placements",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     1,
						Degree:    5,
						RawDegree: 5,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:   "Sun",
						EName:  "Su",
						Rashi:  1,
						Lng:    5,
						Lat:    0,
						Dec:    0,
						Speed:  1,
						Degree: 5,
					},
					{
						Name:   "Moon",
						EName:  "Mo",
						Rashi:  1,
						Lng:    10,
						Lat:    0,
						Dec:    0,
						Speed:  1,
						Degree: 10,
					},
					{
						Name:   "Mars",
						EName:  "Ma",
						Rashi:  1,
						Lng:    20,
						Lat:    0,
						Dec:    0,
						Speed:  1,
						Degree: 20,
					},
				},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if len(v.Grahas) != 3 {
					t.Errorf("expected 3 grahas, got %d", len(v.Grahas))
				}
				for _, g := range v.Grahas {
					if g.BhavaNumber != 1 {
						t.Errorf("expected all grahas in bhava 1, got %d", g.BhavaNumber)
					}
				}
			},
		},
		{
			name: "preserve graha properties in d4 calculation",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     1,
						Degree:    5,
						RawDegree: 5,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:   "Sun",
						EName:  "Su",
						Rashi:  1,
						Lng:    5,
						Lat:    1.5,
						Dec:    0.5,
						Speed:  1.2,
						Degree: 5,
					},
				},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				g := v.Grahas[0]
				if g.Name != "Sun" {
					t.Errorf("expected name Sun, got %s", g.Name)
				}
				if g.EName != "Su" {
					t.Errorf("expected EName Su, got %s", g.EName)
				}
				if g.Lat != 1.5 {
					t.Errorf("expected lat 1.5, got %f", g.Lat)
				}
				if g.Dec != 0.5 {
					t.Errorf("expected dec 0.5, got %f", g.Dec)
				}
				if g.Speed != 1.2 {
					t.Errorf("expected speed 1.2, got %f", g.Speed)
				}
			},
		},
		{
			name: "empty grahas",
			kundli: &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     1,
						Degree:    5,
						RawDegree: 5,
					},
				},
				Grahas: []domain.Graha{},
			},
			wantError: false,
			validate: func(t *testing.T, v domain.Varga) {
				if len(v.Grahas) != 0 {
					t.Errorf("expected 0 grahas, got %d", len(v.Grahas))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d4 := &varga.D4Varga{}
			vargaCalc := &domain.VargaCalc{Kundli: *tt.kundli}

			vargaResult, err := d4.Calculate(vargaCalc)

			if (err != nil) != tt.wantError {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.wantError {
				tt.validate(t, vargaResult)
			}
		})
	}
}

func TestD4VargaPlaceCalculation(t *testing.T) {
	tests := []struct {
		name          string
		degree        float64
		expectedRashi int
	}{
		{name: "place 0 (0-7.5)", degree: 5, expectedRashi: 1},
		{name: "place 1 (7.5-15)", degree: 10, expectedRashi: 5},
		{name: "place 2 (15-22.5)", degree: 20, expectedRashi: 8},
		{name: "place 3 (22.5-30)", degree: 25, expectedRashi: 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kundli := &domain.Kundli{
				Bhavas: []domain.Bhava{
					{
						Number:    1,
						Rashi:     1,
						Degree:    tt.degree,
						RawDegree: tt.degree,
					},
				},
				Grahas: []domain.Graha{
					{
						Name:   "Test",
						EName:  "T",
						Rashi:  1,
						Lng:    tt.degree,
						Lat:    0,
						Dec:    0,
						Speed:  1,
						Degree: tt.degree,
					},
				},
			}

			d4 := &varga.D4Varga{}
			vargaCalc := &domain.VargaCalc{Kundli: *kundli}
			vargaResult, _ := d4.Calculate(vargaCalc)

			if vargaResult.Grahas[0].Rashi != tt.expectedRashi {
				t.Errorf("degree %f: expected rashi %d, got %d", tt.degree, tt.expectedRashi, vargaResult.Grahas[0].Rashi)
			}
		})
	}
}
