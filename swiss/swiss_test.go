package swiss

import (
	"fmt"
	"testing"
	"time"
)

func TestExecSwiss(t *testing.T) {
	dateStr := "13/06/1988"
	timeStr := "12:30"
	lat := -23.5505
	lng := -46.6333

	dateTime, err := time.Parse("02/01/2006 15:04", dateStr+" "+timeStr)
	if err != nil {
		t.Fatalf("Error %v", err)
	}

	options := &SwissOptions{
		DateTime: dateTime,
		Geopos:   LatLng{Lat: lat, Lng: lng},
		Ayanamsa: "1", // Lahiri
	}

	result, err := ExecSwiss(options)

	fmt.Println(result.Bhavas)

	if err != nil {
		t.Fatalf("ExecSwiss exec err: %v. Output: %s", err, result.RawOutput)
	}

	if len(result.Grahas) != 9 {
		t.Error("Invalid Graha Number")
	}

	if len(result.Bhavas) != 12 {
		fmt.Println(len(result.Bhavas))
		t.Error("Invalid Bhava Number")
	}

}
