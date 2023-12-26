package main

import (
	"fmt"
	"math"
	"sync"
)

type CalculationResult struct {
	Radius        float64
	CircleArea    float64
	CircleCircumference float64
	CylinderVolume float64
}

func calculate(radius float64, height float64, resultChannel chan CalculationResult, wg *sync.WaitGroup) {
	defer wg.Done()

	
	circleArea := math.Pi * math.Pow(radius, 2)

	circleCircumference := 2 * math.Pi * radius

	cylinderVolume := circleArea * height

	resultChannel <- CalculationResult{
		Radius:        radius,
		CircleArea:    circleArea,
		CircleCircumference: circleCircumference,
		CylinderVolume: cylinderVolume,
	}
}

func main() {
	radii := []float64{8, 14, 20}
	height := 10

	resultChannel := make(chan CalculationResult, len(radii))

	var wg sync.WaitGroup

	for _, radius := range radii {
		wg.Add(1)
		go calculate(radius, float64(height), resultChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for result := range resultChannel {
		fmt.Printf("Jari-Jari: %.2f\n", result.Radius)
		fmt.Printf("Luas Lingkaran: %.2f\n", result.CircleArea)
		fmt.Printf("Keliling Lingkaran: %.2f\n", result.CircleCircumference)
		fmt.Printf("Volume Tabung: %.2f\n", result.CylinderVolume)
		fmt.Println()
	}
}
