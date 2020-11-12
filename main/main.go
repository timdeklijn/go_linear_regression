package main

import (
	"fmt"
)

func main() {
	// Data
	length := Vec{5, 5.5, 5.8, 6, 6.3, 6.5, 6.9, 7.1, 7.4, 7.7}
	weight := Vec{7, 7.5, 7.3, 8, 7.7, 7.9, 8.2, 8.8, 8.4, 8.3}

	// Setup
	x, _ := addOnes(length)
	lr := NewLinearRegression(x, weight)
	lr.config.epochs = 10000
	// Run
	lr.Fit()
	fmt.Println("\nThetas:", lr.thetas)

	var sample Arr = Arr{Vec{1, 10}}
	fmt.Println("Prediction:", lr.Predict(sample))
}
