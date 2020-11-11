package main

import (
	"reflect"
	"testing"
)

func TestTransposeMat(t *testing.T) {
	ts := []struct {
		test [][]float32
		want [][]float32
	}{{
		[][]float32{[]float32{1}, []float32{2}},
		[][]float32{[]float32{1, 2}},
	}, {
		[][]float32{[]float32{1, 2}, []float32{3, 4}},
		[][]float32{[]float32{1, 3}, []float32{2, 4}},
	},
	}

	for _, tc := range ts {
		got := transposeMat(tc.test)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TransposeMat was incorrect, got %v, want: %v", got, tc.want)
		}
	}
}

func TestTransposeVec(t *testing.T) {
	ts := []struct {
		test []float32
		want [][]float32
	}{{
		[]float32{1},
		[][]float32{[]float32{1}},
	}, {
		[]float32{1, 2},
		[][]float32{[]float32{1}, []float32{2}},
	},
	}

	for _, tc := range ts {
		got := transposeVec(tc.test)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TransposeVec was incorrect, got %v, want: %v", got, tc.want)
		}
	}
}
