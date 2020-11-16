package algebra

import (
	"reflect"
	"testing"
)

func TestVec_Add(t *testing.T) {
	type args struct {
		other Vec
	}
	tests := []struct {
		name string
		v    Vec
		args args
		want Vec
	}{
		{
			name: "Simple add test",
			v:    Vec{2, 3, 4},
			args: args{
				other: Vec{1, 2, 3},
			},
			want: Vec{3, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Add(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec_AddOnes(t *testing.T) {
	tests := []struct {
		name string
		v    Vec
		want Arr
	}{
		{
			name: "Test add ones for vector",
			v:    Vec{1, 2},
			want: Arr{Vec{1, 1}, Vec{1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.AddOnes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec_Sub(t *testing.T) {
	type args struct {
		other Vec
	}
	tests := []struct {
		name string
		v    Vec
		args args
		want Vec
	}{
		{
			name: "Simple sub test",
			v:    Vec{2, 3, 4},
			args: args{
				other: Vec{1, 2, 3},
			},
			want: Vec{1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Sub(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec_Transpose(t *testing.T) {
	tests := []struct {
		name string
		v    Vec
		want Arr
	}{
		{
			name: "transpose 1 element vector",
			v:    Vec{1},
			want: Arr{Vec{1}},
		},
		{
			name: "transpose 2 element vector",
			v:    Vec{1, 2},
			want: Arr{Vec{1}, Vec{2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Transpose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
