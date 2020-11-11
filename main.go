package main

import (
	"fmt"
	"math/rand"
)

func addOnes(a []float32) [][]float32 {
	var new [][]float32
	for _, aa := range a {
		new = append(new, []float32{1.0, aa})
	}
	return new
}

func predict(ll [][]float32, thetas []float32) []float32 {
	var predictions []float32
	for _, i := range ll {
		predictions = append(predictions, i[0]*thetas[0]+i[1]*thetas[1])
	}
	return predictions
}

func transposeMat(m [][]float32) [][]float32 {
	var new [][]float32
	for i := 0; i < len(m[0]); i++ {
		var row []float32
		new = append(new, row)
		for j := 0; j < len(m); j++ {
			new[i] = append(new[i], m[j][i])
		}
	}
	return new
}

func transposeVec(v []float32) [][]float32 {
	var new [][]float32
	for _, i := range v {
		new = append(new, []float32{i})
	}
	return new
}

func calcGradient(ll [][]float32, residuals []float32, lr float32, m int) []float32 {
	var new []float32
	x := transposeMat(ll)
	r := transposeVec(residuals)
	// i => 0, 1
	for i := 0; i < len(x); i++ {
		// j => 0,1,2,3,4,5,6,7,8,9
		var n float32 = 0.0
		for j := 0; j < len(r); j++ {
			n += x[i][j] * r[j][0]
		}
		new = append(new, n)

	}
	s := lr / float32(m)
	for i := range new {
		new[i] *= s
	}
	return new
}

func calcError(t []float32, l [][]float32, w []float32) float32 {
	p := predict(l, t)
	var s float32 = 0.0
	for i := range p {
		d := w[i] - p[i]
		s += d * d
	}
	return s
}

func main() {
	// length := []float32{5, 5.5, 5.8, 6, 6.3, 6.5, 6.9, 7.1, 7.4, 7.7}
	// weight := []float32{7, 7.5, 7.3, 8, 7.7, 7.9, 8.2, 8.8, 8.4, 8.3}
	// y = 2x + 1 -> y = ax + b -> weights [2,1]
	length := []float32{1, 2, 3, 4}
	weight := []float32{3, 5, 7, 9}
	ll := addOnes(length)
	var thetas []float32
	for i := 0; i < len(ll[0]); i++ {
		thetas = append(thetas, rand.Float32())
	}
	cntr := 0
	for cntr < 100 {
		pred := predict(ll, thetas)
		var residuals []float32
		for i := range pred {
			residuals = append(residuals, weight[i]-pred[i])
		}
		g := calcGradient(ll, residuals, 0.1, len(length))
		for i := range thetas {
			thetas[i] += g[i]
		}

		e := calcError(thetas, ll, weight)
		fmt.Println(thetas, e, pred)
		cntr++
	}
}
