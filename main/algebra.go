package main

import "fmt"

// Vec (tor) type
type Vec []float32

// Arr (ay) type
type Arr []Vec

func addOnes(a interface{}) (Arr, error) {
	var new Arr
	switch v := a.(type) {
	case Arr:
		for _, ii := range v {
			var tmp Vec
			tmp = append(tmp, 1.0)
			for _, iii := range ii {
				tmp = append(tmp, iii)
			}
			new = append(new, tmp)
		}
		return new, nil
	case Vec:
		for _, aa := range v {
			new = append(new, Vec{1.0, aa})
		}
		return new, nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("This function only accepts Arr and Vec types, not '%v'", a))
	}
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
