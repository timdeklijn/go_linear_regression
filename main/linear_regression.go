package main

import (
	"fmt"
	"math/rand"
)

// LinearRegressionConfig is a container for settings for running
// linear regression.
type linearRegressionConfig struct {
	learningRate   float32
	printFrequency int
	epochs         int
}

// LinearRegressionData is a container for the data that linear regression
// will be run on.
type linearRegressionData struct {
	x Arr
	y Vec
	l int // length of the x data
}

// LinearRegression is the type bundling the linear regression containers and
// functions
type LinearRegression struct {
	config linearRegressionConfig
	data   linearRegressionData
	thetas Vec
}

// NewLinearRegression creates a new LinearRegression type with sane defaults
func NewLinearRegression(x Arr, y Vec) LinearRegression {
	config := linearRegressionConfig{
		learningRate:   0.01,
		printFrequency: 100,
		epochs:         1000,
	}
	data := linearRegressionData{x, y, len(x)}
	thetas := initThetas(len(x[0]))
	return LinearRegression{config, data, thetas}
}

// Predict calculates the values calculated from the x and thetas by
// taking the dot product.
//
//		y = x.thetas
func (lr *LinearRegression) Predict() Vec {
	var predictions Vec
	for _, i := range lr.data.x {
		predictions = append(predictions, i[0]*lr.thetas[0]+i[1]*lr.thetas[1])
	}
	return predictions
}

// Fit runs the least squares optimization to obtain the thetas
// the best fit the data.
func (lr *LinearRegression) Fit() {
	cntr := 1
	pred := lr.Predict()
	for cntr < lr.config.epochs+1 {
		var residuals Vec
		for i := range pred {
			residuals = append(residuals, lr.data.y[i]-pred[i])
		}
		g := lr.calcGradient(residuals)
		for i := range lr.thetas {
			lr.thetas[i] += g[i]
		}

		pred := lr.Predict()
		if cntr%100 == 0 {
			e := lr.calcError(pred)
			fmt.Println(
				"================================",
				"\nEpoch     ", cntr,
				"\nError     ", e,
				"\nThetas    ", lr.thetas,
			)
		}
		cntr++
	}
}

func (lr *LinearRegression) calcGradient(residuals Vec) Vec {
	var new Vec
	x := transposeMat(lr.data.x)
	r := transposeVec(residuals)
	for i := 0; i < len(x); i++ {
		var n float32 = 0.0
		for j := 0; j < len(r); j++ {
			n += x[i][j] * r[j][0]
		}
		new = append(new, n)
	}
	s := lr.config.learningRate / float32(lr.data.l)
	for i := range new {
		new[i] *= s
	}
	return new
}

// calcError calculates the squared error between a predictions and
// the truth value:
//
//		e = (y - y_pred)**2
func (lr *LinearRegression) calcError(p Vec) float32 {
	var s float32 = 0.0
	for i := range p {
		d := lr.data.y[i] - p[i]
		s += d * d
	}
	return s
}

func initThetas(n int) Vec {
	var t Vec
	for i := 0; i < n; i++ {
		t = append(t, rand.Float32())
	}
	fmt.Println(t)
	return t
}
