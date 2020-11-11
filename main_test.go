package main

import (
	"reflect"
	"testing"
)

func TestTransposeMat(t *testing.T) {
	ts := []struct {
		test arr
		want arr
	}{{
		arr{vec{1}, vec{2}},
		arr{vec{1, 2}},
	}, {
		arr{vec{1, 2}, vec{3, 4}},
		arr{vec{1, 3}, vec{2, 4}},
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
		test vec
		want arr
	}{{
		vec{1},
		arr{vec{1}},
	}, {
		vec{1, 2},
		arr{vec{1}, vec{2}},
	},
	}

	for _, tc := range ts {
		got := transposeVec(tc.test)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TransposeVec was incorrect, got %v, want: %v", got, tc.want)
		}
	}
}
