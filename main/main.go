package main

import (
	"fmt"
)

// Vec (tor) type
type Vec []float32

// Arr (ay) type
type Arr []Vec

func addOnes(a Vec) Arr {
	var new Arr
	for _, aa := range a {
		new = append(new, Vec{1.0, aa})
	}
	return new
}

func transposeMat(m Arr) Arr {
	var new Arr
	for i := 0; i < len(m[0]); i++ {
		var row Vec
		new = append(new, row)
		for j := 0; j < len(m); j++ {
			new[i] = append(new[i], m[j][i])
		}
	}
	return new
}

func transposeVec(v Vec) Arr {
	var new Arr
	for _, i := range v {
		new = append(new, Vec{i})
	}
	return new
}

func main() {
	// Data
	length := Vec{5, 5.5, 5.8, 6, 6.3, 6.5, 6.9, 7.1, 7.4, 7.7}
	weight := Vec{7, 7.5, 7.3, 8, 7.7, 7.9, 8.2, 8.8, 8.4, 8.3}

	// Setup
	lr := NewLinearRegression(addOnes(length), weight)
	lr.config.epochs = 10000
	// Run
	lr.Fit()
	fmt.Println("\nThetas:", lr.thetas)

	var sample Arr = Arr{Vec{1, 10}}
	fmt.Println("Prediction:", lr.Predict(sample))
}
