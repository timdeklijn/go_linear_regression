package main

import (
	"reflect"
	"testing"
)

func TestLinearRegression_Fit(t *testing.T) {
	type fields struct {
		config LinearRegressionConfig
		data   LinearRegressionData
		thetas Vec
	}
	tests := []struct {
		name   string
		fields fields
		want   Vec
	}{
		{
			name: "Simple Test Fit: y=a+bx (a=0,b=1)",
			fields: fields{
				LinearRegressionConfig{
					learningRate:   0.001,
					printUpdates:   false,
					printFrequency: 100,
					epochs:         100000,
				},
				LinearRegressionData{
					x: Vec{1, 2, 3, 4}.AddOnes(),
					y: Vec{1, 2, 3, 4},
					l: 4,
				},
				InitThetas(2),
			},
			want: Vec{0.0, 1.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &LinearRegression{
				config: tt.fields.config,
				data:   tt.fields.data,
				thetas: tt.fields.thetas,
			}
			// Run the actual fit
			lr.Fit()
			// And check the results against the expected values
			got := lr.thetas
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fit() = %v, want %v", got, tt.want)
			}

		})
	}
}
