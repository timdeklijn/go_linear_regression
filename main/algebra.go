package main

// Vec (tor) type
type Vec []float32

// Arr (ay) type
type Arr []Vec

// AddOnes creates an Array where each element is a list (Vec)
// starting with 1.0.
func (a Arr) AddOnes() Arr {
	var new Arr
	for i := 0; i < len(a); i++ {
		var tmp Vec
		tmp = append(tmp, 1.0)
		for _, ii := range a[i] {
			tmp = append(tmp, ii)
		}
		new = append(new, tmp)

	}
	return new
}

// AddOnes creates an Array where each element is a list (Vec)
// starting with 1.0.
func (v Vec) AddOnes() Arr {
	var new Arr
	for _, aa := range v {
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
