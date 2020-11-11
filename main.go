package main

import (
	"fmt"
	"math/rand"
)

type vec []float32
type arr []vec

func addOnes(a vec) arr {
	var new arr
	for _, aa := range a {
		new = append(new, vec{1.0, aa})
	}
	return new
}

func predict(ll arr, thetas vec) vec {
	var predictions vec
	for _, i := range ll {
		predictions = append(predictions, i[0]*thetas[0]+i[1]*thetas[1])
	}
	return predictions
}

func transposeMat(m arr) arr {
	var new arr
	for i := 0; i < len(m[0]); i++ {
		var row vec
		new = append(new, row)
		for j := 0; j < len(m); j++ {
			new[i] = append(new[i], m[j][i])
		}
	}
	return new
}

func transposeVec(v vec) arr {
	var new arr
	for _, i := range v {
		new = append(new, vec{i})
	}
	return new
}

func calcGradient(ll arr, residuals vec, lr float32, m int) vec {
	var new vec
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

func calcError(t vec, l arr, w vec) float32 {
	p := predict(l, t)
	var s float32 = 0.0
	for i := range p {
		d := w[i] - p[i]
		s += d * d
	}
	return s
}

func main() {
	length := vec{5, 5.5, 5.8, 6, 6.3, 6.5, 6.9, 7.1, 7.4, 7.7}
	weight := vec{7, 7.5, 7.3, 8, 7.7, 7.9, 8.2, 8.8, 8.4, 8.3}
	// y = 2x + 1 -> y = ax + b -> weights [2,1]
	// length := vec{1, 2, 3, 4}
	// weight := vec{3, 5, 7, 9}
	ll := addOnes(length)
	var thetas vec
	for i := 0; i < len(ll[0]); i++ {
		thetas = append(thetas, rand.Float32())
	}
	cntr := 1
	for cntr < 100001 {
		pred := predict(ll, thetas)
		var residuals vec
		for i := range pred {
			residuals = append(residuals, weight[i]-pred[i])
		}
		g := calcGradient(ll, residuals, 0.01, len(length))
		for i := range thetas {
			thetas[i] += g[i]
		}

		if cntr%100 == 0 {
			e := calcError(thetas, ll, weight)
			fmt.Println("Epoch", cntr, "\nThetas", thetas, "\nError", e, "\nPrediction", pred)
			fmt.Println("=======================")
		}
		cntr++
	}
}
