package main

import (
	"reflect"
	"testing"
)

func TestTransposeMat(t *testing.T) {
	ts := []struct {
		test Arr
		want Arr
	}{{
		Arr{Vec{1}, Vec{2}},
		Arr{Vec{1, 2}},
	}, {
		Arr{Vec{1, 2}, Vec{3, 4}},
		Arr{Vec{1, 3}, Vec{2, 4}},
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
		test Vec
		want Arr
	}{{
		Vec{1},
		Arr{Vec{1}},
	}, {
		Vec{1, 2},
		Arr{Vec{1}, Vec{2}},
	},
	}

	for _, tc := range ts {
		got := transposeVec(tc.test)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TransposeVec was incorrect, got %v, want: %v", got, tc.want)
		}
	}
}
