package algebra

import (
	"reflect"
	"testing"
)

func TestArr_AddOnes(t *testing.T) {
	tests := []struct {
		name string
		a    Arr
		want Arr
	}{
		{
			name: "AddOnes to simple array",
			a:    Arr{Vec{1, 2}, Vec{1, 2}},
			want: Arr{Vec{1, 1, 2}, Vec{1, 1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AddOnes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArr_Transpose(t *testing.T) {
	tests := []struct {
		name string
		a    Arr
		want Arr
	}{
		{
			name: "Transpose (1,2) array",
			a:    Arr{Vec{1}, Vec{2}},
			want: Arr{Vec{1, 2}},
		},
		{
			name: "Transpose (2,2) array",
			a:    Arr{Vec{1, 2}, Vec{3, 4}},
			want: Arr{Vec{1, 3}, Vec{2, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Transpose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
