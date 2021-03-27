package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lytics/anomalyzer"
)

func randomA(min, max int) int {
	return rand.Intn(max-min) + min
}

func anomaly(MAX int) {

	conf := &anomalyzer.AnomalyzerConf{
		Sensitivity: 0.1,
		UpperBound:  5,
		LowerBound:  anomalyzer.NA,
		ActiveSize:  1,
		NSeasons:    4,
		Methods:     []string{"diff", "fence", "magnitude", "ks"},
	}

	data := []float64{}
	SEED := time.Now().Unix()
	rand.Seed(SEED)

	for i := 0; i < MAX; i++ {
		data = append(data, float64(randomA(0, MAX)))
	}
	fmt.Println("data:", data)

	anom, _ := anomalyzer.NewAnomalyzer(conf, data)
	prob := anom.Push(8.0)
	fmt.Println("Anomalous Probability:", prob)
}
