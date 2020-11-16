package algebra

// Vec (tor) type
type Vec []float32

// AddOnes creates an Array where each element is a list (Vec)
// starting with 1.0.
func (v Vec) AddOnes() Arr {
	var arr Arr
	for _, aa := range v {
		arr = append(arr, Vec{1.0, aa})
	}
	return arr
}

// Transpose returns the transpose of the vector,
// which is an Array.
func (v Vec) Transpose() Arr {
	var arr Arr
	for _, i := range v {
		arr = append(arr, Vec{i})
	}
	return arr
}

// Sub subtracts other from v. Returns a new Vec.
func (v Vec) Sub(other Vec) Vec {
	var vec Vec
	for i := range v {
		vec = append(vec, v[i]-other[i])
	}
	return vec
}

// Add add other to v. Returns a new Vec.
func (v Vec) Add(other Vec) Vec {
	var vec Vec
	for i := range v {
		vec = append(vec, v[i]+other[i])
	}
	return vec
}
