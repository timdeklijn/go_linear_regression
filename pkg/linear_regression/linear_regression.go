package linear_regression

import (
	"fmt"
	. "github.com/timdeklijn/go_linear_regression/pkg/algebra"
	"math/rand"
)

// LinearRegressionConfig is a container for settings for running
// linear regression.
type LinearRegressionConfig struct {
	learningRate   float32 // How big the changes in theta will be
	printUpdates   bool    // print updates at printFrequency
	printFrequency int     // How many epochs should there be a print
	epochs         int     // Number of update cycles
}

// LinearRegressionData is a container for the data that linear regression
// will be run on.
type LinearRegressionData struct {
	x Arr // Features
	y Vec // Labels
	l int // length of the x data
}

// LinearRegression is the type bundling the linear regression containers and
// functions
type LinearRegression struct {
	config LinearRegressionConfig // Config of the linear regression
	data   LinearRegressionData   // Data container
	Thetas Vec                    // Thetas, will be optimized
}

// NewLinearRegression creates a new LinearRegression type with sane defaults
func NewLinearRegression(x Arr, y Vec) LinearRegression {
	config := LinearRegressionConfig{
		learningRate:   0.01,
		printUpdates:   true,
		printFrequency: 1000,
		epochs:         100000,
	}
	data := LinearRegressionData{x, y, len(x)}
	thetas := InitThetas(len(x[0]))
	return LinearRegression{config, data, thetas}
}

// Predict calculates the values calculated from the x and thetas by
// taking the dot product.
//
//		y = x.thetas
func (lr *LinearRegression) Predict(d Arr) Vec {
	return d.DotVec(lr.Thetas)
}

// calcResiduals calculates the difference between the predicted values,
// pred, and the actual values lr.data.y. The difference is returned as
// a vector
func (lr *LinearRegression) calcResiduals(pred Vec) Vec {
	return lr.data.y.Sub(pred)
}

// Fit runs the least squares optimization to obtain the thetas
// the best fit the data.
func (lr *LinearRegression) Fit() {
	cntr := 1
	pred := lr.Predict(lr.data.x)
	for cntr < lr.config.epochs+1 {
		residuals := lr.calcResiduals(pred)
		gradient := lr.calcGradient(residuals)
		lr.Thetas = lr.Thetas.Add(gradient)
		// Create new predictions
		pred = lr.Predict(lr.data.x)
		if cntr%100 == 0 && lr.config.printUpdates {
			e := lr.calcError(pred)
			fmt.Println(
				"================================",
				"\nEpoch     ", cntr,
				"\nError     ", e,
				"\nThetas    ", lr.Thetas,
			)
		}
		cntr++
	}
}

// calcGradient calculates:
//
//		(self.lr/m)*(np.dot(X.T, residuals.T))
//
func (lr *LinearRegression) calcGradient(residuals Vec) Vec {
	var grad Vec
	x := lr.data.x.Transpose()
	r := residuals.Transpose()
	for i := 0; i < len(x); i++ {
		var n float32 = 0.0
		for j := 0; j < len(r); j++ {
			n += x[i][j] * r[j][0]
		}
		grad = append(grad, n)
	}
	s := lr.config.learningRate / float32(lr.data.l)
	for i := range grad {
		grad[i] *= s
	}
	return grad
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

// InitThetas creates a vector with random numbers of length n. This will
// be used as the starting point for the thetas that will be optimized during
// linear regression.
func InitThetas(n int) Vec {
	var t Vec
	for i := 0; i < n; i++ {
		t = append(t, rand.Float32())
	}
	return t
}
