package algebra

// Arr (ay) type
type Arr []Vec

// AddOnes creates an Array where each element is a list (Vec)
// starting with 1.0.
func (a Arr) AddOnes() Arr {
	var arr Arr
	for i := 0; i < len(a); i++ {
		var tmp Vec
		tmp = append(tmp, 1.0)
		for _, ii := range a[i] {
			tmp = append(tmp, ii)
		}
		arr = append(arr, tmp)

	}
	return arr
}

// Transpose returns a transpose of the Array.
func (a Arr) Transpose() Arr {
	var arr Arr
	for i := 0; i < len(a[0]); i++ {
		var row Vec
		arr = append(arr, row)
		for j := 0; j < len(a); j++ {
			arr[i] = append(arr[i], a[j][i])
		}
	}
	return arr
}

// DotVec calculates the dotproduct between the array and
// input vector v.
func (a Arr) DotVec(v Vec) Vec {
	// TODO: add size check and return error
	var vec Vec
	for i := range a {
		var tmpSum float32 = 0.0
		for j := range v {
			tmpSum += a[i][j] * v[j]
		}
		vec = append(vec, tmpSum)
	}
	return vec
}

// DotArr calculates the dotproduct between two arrays
func (a Arr) DotArr(other Arr) Vec {
	// TODO: Make work for multi column other
	var vec Vec
	for c := 0; c < len(other[0]); c++ {
		for i := range a {
			var tmpSum float32 = 0.0
			for j := range a[0] {
				tmpSum += a[i][j] * other[j][c]
			}
			vec = append(vec, tmpSum)
		}
	}
	return vec
}
