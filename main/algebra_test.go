package main

import (
	"fmt"
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

func TestAddOnes(t *testing.T) {
	cases := []struct {
		test interface{}
		want Arr
		Err  error
	}{
		{
			Vec{1, 2},
			Arr{Vec{1, 1}, Vec{1, 2}},
			nil,
		},
		{
			Arr{Vec{1, 2}, Vec{1, 2}},
			Arr{Vec{1, 1, 2}, Vec{1, 1, 2}},
			nil,
		},
		{
			"Wrong Type",
			nil,
			fmt.Errorf("This function only accepts Arr and Vec types, not 'Wrong Type'"),
		},
	}
	for _, tc := range cases {
		got, err := addOnes(tc.test)
		fmt.Println(got)
		if err != tc.Err {
			fmt.Println(tc.Err)
			t.Errorf("Unexpected error: %s", err)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("AddOnes was incorrect, got %v, want: %v", got, tc.want)
		}
	}
}
