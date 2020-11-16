package main

import (
	"fmt"
	"linear_regression/pkg/algebra"
	. "linear_regression/pkg/linear_regression"
)

func main() {
	// Data
	length := algebra.Vec{5, 5.5, 5.8, 6, 6.3, 6.5, 6.9, 7.1, 7.4, 7.7}
	weight := algebra.Vec{7, 7.5, 7.3, 8, 7.7, 7.9, 8.2, 8.8, 8.4, 8.3}

	// Setup
	x := length.AddOnes()
	lr := NewLinearRegression(x, weight)
	// Run
	lr.Fit()
	fmt.Println("\nThetas:", lr.Thetas)

	var sample = algebra.Arr{algebra.Vec{1, 10}}
	fmt.Println("Prediction:", lr.Predict(sample))
}
