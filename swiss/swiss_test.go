package swiss

import (
	"fmt"
	"testing"
	"time"
)

func TestExecSwiss(t *testing.T) {
	dateStr := "13/06/1988"
	timeStr := "09:30"
	lat := -23.5505
	lng := -46.6333

	dateTime, err := time.Parse("02/01/2006 15:04", dateStr+" "+timeStr)
	if err != nil {
		t.Fatalf("Erro ao fazer parse da data e hora: %v", err)
	}

	// Opções para o swetest
	options := &SwissOptions{
		DateTime: dateTime,
		Geopos:   LatLng{Lat: lat, Lng: lng},
		Ayanamsa: "1", // Lahiri
	}

	// Executar o swetest
	result, err := ExecSwiss(options)
	fmt.Println(result)

	if err != nil {
		t.Fatalf("ExecSwiss retornou um erro: %v. Output: %s", err, result.RawOutput)
	}

	// Verificar se o resultado contém planetas
	if len(result.Grahas) == 0 {
		t.Error("Nenhum graha (planeta) foi encontrado no resultado.")
	}

	// Verificar se o resultado contém casas
	if len(result.Bhavas) == 0 {
		t.Error("Nenhum bhava (casa) foi encontrado no resultado.")
	}

	// Opcional: Verificar um planeta específico (ex: Sol)
	foundSun := false
	for _, graha := range result.Grahas {
		if graha.Name == "Sun" {
			foundSun = true
			t.Logf("Sol encontrado em: %s", graha.Longitude)
			break
		}
	}
	if !foundSun {
		t.Error("Sol não foi encontrado na lista de grahas.")
	}
}
