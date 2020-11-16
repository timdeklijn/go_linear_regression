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
		got := tc.test.Transpose()
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
		got := tc.test.Transpose()
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TransposeVec was incorrect, got %v, want: %v", got, tc.want)
		}
	}
}

func TestAddOnesArr(t *testing.T) {
	cases := []struct {
		test Arr
		want Arr
	}{

		{
			Arr{Vec{1, 2}, Vec{1, 2}},
			Arr{Vec{1, 1, 2}, Vec{1, 1, 2}},
		},
	}
	for _, tc := range cases {
		got := tc.test.AddOnes()
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("AddOnes was incorrect, got %v, want: %v", got, tc.want)
		}
	}
}

func TestAddOnesVec(t *testing.T) {
	cases := []struct {
		test Vec
		want Arr
	}{

		{
			Vec{1, 2},
			Arr{Vec{1, 1}, Vec{1, 2}},
		},
	}
	for _, tc := range cases {
		got := tc.test.AddOnes()
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("AddOnes was incorrect, got %v, want: %v", got, tc.want)
		}
	}
}

func TestSubVec(t *testing.T) {
	cases := []struct {
		test    Vec
		testMin Vec
		want    Vec
	}{{
		Vec{2, 3, 4},
		Vec{1, 2, 3},
		Vec{1, 1, 1},
	}}
	for _, tc := range cases {
		got := tc.test.Sub(tc.testMin)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Sub was incorrect, got %v, want: %v", got, tc.want)

		}
	}
}

func TestAddVec(t *testing.T) {
	cases := []struct {
		test    Vec
		testMin Vec
		want    Vec
	}{{
		Vec{2, 3, 4},
		Vec{1, 2, 3},
		Vec{3, 5, 7},
	}}
	for _, tc := range cases {
		got := tc.test.Add(tc.testMin)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Sub was incorrect, got %v, want: %v", got, tc.want)

		}
	}
}
