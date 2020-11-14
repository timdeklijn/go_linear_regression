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

// Transpose eturns a transpose of the Array.
func (a Arr) Transpose() Arr {
	var new Arr
	for i := 0; i < len(a[0]); i++ {
		var row Vec
		new = append(new, row)
		for j := 0; j < len(a); j++ {
			new[i] = append(new[i], a[j][i])
		}
	}
	return new
}

// Transpose returns the transpose of the vector,
// which is an Array.
func (v Vec) Transpose() Arr {
	var new Arr
	for _, i := range v {
		new = append(new, Vec{i})
	}
	return new
}

// Sub subtracts other from v. Returns a new Vec.
func (v Vec) Sub(other Vec) Vec {
	var new Vec
	for i := range v {
		new = append(new, v[i]-other[i])
	}
	return new
}
