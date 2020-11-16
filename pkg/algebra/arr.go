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
