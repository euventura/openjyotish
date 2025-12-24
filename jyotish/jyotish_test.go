package jyotish

import (
	"fmt"
	"testing"
)

func TestGrahaBuild(t *testing.T) {

	graha := Graha{
		Name:  "Sun",
		Lat:   2.0,
		Lng:   59.0,
		Dec:   0.5,
		Speed: 0.2,
	}

	grahaBuild(&graha, Bhava{Number: 1, RawDegree: 111})

	if graha.Rashi != 2 {
		t.Errorf("Expected Rashi to be 2, but got %d", graha.Rashi)
	}

	if graha.BhavaNumber != 11 {
		t.Errorf("Expected BhavaNumber to be 11, but got %d", graha.BhavaNumber)
	}

	graha = Graha{
		Name:  "Sun",
		Lat:   2.0,
		Lng:   121.0,
		Dec:   0.5,
		Speed: 0.2,
	}

	grahaBuild(&graha, Bhava{Number: 1, RawDegree: 0})
	fmt.Println(graha.BhavaNumber)

	if graha.BhavaNumber != 5 {
		t.Errorf("Expected BhavaNumber to be 5, but got %d", graha.BhavaNumber)
	}

}
