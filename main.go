package main

import (
	"fmt"
	"log"
	"openjyotish/swiss"
	"time"
)

func main() {
	// Dados fornecidos pelo usuário
	dateStr := "13/06/1988"
	timeStr := "09:30"
	lat := -23.5505
	lng := -46.6333

	// Parsing da data e hora
	dateTime, err := time.Parse("02/01/2006 15:04", dateStr+" "+timeStr)
	if err != nil {
		log.Fatalf("Erro ao fazer parse da data e hora: %v", err)
	}

	// Opções para o swetest
	options := &swiss.SwissOptions{
		DateTime: dateTime,
		Geopos:   swiss.LatLng{Lat: lat, Lng: lng},
		Ayanamsa: "1", // Lahiri
	}

	// Executar o swetest
	result, err := swiss.ExecSwiss(options)
	if err != nil {
		log.Fatalf("ExecSwiss retornou um erro: %v", err)
	}

	// Imprimir o resultado
	fmt.Println("Planetas:")
	for _, graha := range result.Grahas {
		fmt.Printf("- %s: %s\n", graha.Name, graha.Longitude)
	}

	fmt.Println("\nCasas (Bhavas):")
	for _, bhava := range result.Bhavas {
		fmt.Printf("- Casa %d: Início %f, Fim %f\n", bhava.HouseNumber, bhava.Start, bhava.End)
	}
}