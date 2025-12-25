package test

import (
	"testing"
	"time"

	"openjyotish/internal/application"
	"openjyotish/internal/domain"
	"openjyotish/swiss"
)

func TestNakshatraService_CalcNakshatra(t *testing.T) {
	ns := application.NakshatraService{}
	nakshatra, err := ns.CalcNakshatra(10.0)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if nakshatra.Name != "Ashwini" {
		t.Fatalf("expected Ashwini, got %s", nakshatra.Name)
	}
	if nakshatra.Pada != 4 {
		t.Fatalf("expected pada 4, got %d", nakshatra.Pada)
	}
}

func TestNakshatraService_CalcNakshatra_Invalid(t *testing.T) {
	ns := application.NakshatraService{}
	_, err := ns.CalcNakshatra(400.0)
	if err == nil {
		t.Fatalf("expected error for invalid degree")
	}
}

func TestDashaService_CalculateDashas(t *testing.T) {
	ds := application.DashaService{}
	k := domain.Kundli{
		Grahas: []domain.Graha{
			{EName: "Moon", Lng: 10.0},
			{EName: "Lagna", Lng: 0.0},
		},
	}
	birthDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	if err := ds.CalculateDashas(&k, birthDate); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	dasha, ok := k.Dasha["Vimsottari"]
	if !ok {
		t.Fatalf("expected Vimsottari in dasha map")
	}
	if len(dasha.Periods) == 0 {
		t.Fatalf("expected periods to be populated")
	}
}

func TestKundliService_LoadFromSwiss(t *testing.T) {
	ks := application.NewKundliService(&application.NakshatraService{}, &application.DashaService{})
	swissRes := &swiss.Result{
		Bhavas: []swiss.Bhava{
			{Number: 1, Start: 0.0, End: 30.0},
			{Number: 2, Start: 30.0, End: 60.0},
		},
		Grahas: []swiss.Graha{
			{Name: "Moon", Longitude: 10.0, Latitude: 0.0, Dec: 0.0, Speed: 0.0},
			{Name: "Lagna", Longitude: 0.0, Latitude: 0.0, Dec: 0.0, Speed: 0.0},
		},
	}
	birthDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	k := ks.LoadFromSwiss(swissRes, birthDate)

	if len(k.Grahas) != 2 {
		t.Fatalf("expected 2 grahas, got %d", len(k.Grahas))
	}

	moon := findGrahaByEName(k.Grahas, "Moon")
	if moon == nil {
		t.Fatalf("expected Moon graha")
	}
	if moon.Name != "Chandra" {
		t.Fatalf("expected Chandra mapped name, got %s", moon.Name)
	}
	if moon.Nakshatra.Name == "" {
		t.Fatalf("expected nakshatra to be filled")
	}

	if _, ok := k.Dasha["Vimsottari"]; !ok {
		t.Fatalf("expected Vimsottari dasha computed")
	}
}

func findGrahaByEName(grahas []domain.Graha, name string) *domain.Graha {
	for i := range grahas {
		if grahas[i].EName == name {
			return &grahas[i]
		}
	}
	return nil
}
