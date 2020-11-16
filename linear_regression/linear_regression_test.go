package main

import (
	"math"
	"reflect"
	"testing"
)

func almostEqual(a, b, epsilon float32) bool {
	return float32(math.Abs(float64(a-b))) <= epsilon
}

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
			for i := range got {
				// With margin, check if thetas are equal to want
				if !almostEqual(got[i], tt.want[i], 1e-3) {
					t.Errorf("Fit() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestLinearRegression_Predict(t *testing.T) {
	type fields struct {
		config LinearRegressionConfig
		data   LinearRegressionData
		thetas Vec
	}
	type args struct {
		d Arr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec
	}{
		{
			name: "Predict simple: y(x)=1+x->x=2->y=3",
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
				Vec{1, 1},
			},
			args: args{
				d: Vec{2}.AddOnes(),
			},
			want: Vec{3},
		},
		{
			name: "Predict simpler: y(x)=0+x->x=2->y=2",
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
				Vec{0, 1},
			},
			args: args{
				d: Vec{2}.AddOnes(),
			},
			want: Vec{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &LinearRegression{
				config: tt.fields.config,
				data:   tt.fields.data,
				thetas: tt.fields.thetas,
			}
			if got := lr.Predict(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Predict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinearRegression_calcResiduals(t *testing.T) {
	type fields struct {
		config LinearRegressionConfig
		data   LinearRegressionData
		thetas Vec
	}
	type args struct {
		pred Vec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec
	}{
		{
			name: "Return a 0 residual.",
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
				Vec{1, 1},
			},
			args: args{
				pred: Vec{1, 2, 3, 4},
			},
			want: Vec{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &LinearRegression{
				config: tt.fields.config,
				data:   tt.fields.data,
				thetas: tt.fields.thetas,
			}
			if got := lr.calcResiduals(tt.args.pred); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcResiduals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinearRegression_calcGradient(t *testing.T) {
	type fields struct {
		config LinearRegressionConfig
		data   LinearRegressionData
		thetas Vec
	}
	type args struct {
		residuals Vec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec
	}{
		{
			// X.T = [[1,1],[1,2],[1,3],[1,4]].T = [[1,1,1,1],[1,2,3,4]]
			// residuals.T [1,1,1,1].T = [[1],[1],[1],[1]]
			// X.T . residuals.T = [[4],[10]]
			// scale = lr / m = 2.5
			// grad = scale * X.T . residuals.T =  [10,25]
			name: "Calculate the gradient.",
			fields: fields{
				LinearRegressionConfig{
					learningRate:   10,
					printUpdates:   false,
					printFrequency: 100,
					epochs:         100000,
				},
				LinearRegressionData{
					x: Vec{1, 2, 3, 4}.AddOnes(),
					y: Vec{1, 2, 3, 4},
					l: 4,
				},
				Vec{1, 1},
			},
			args: args{
				residuals: Vec{1, 1, 1, 1},
			},
			want: Vec{10, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &LinearRegression{
				config: tt.fields.config,
				data:   tt.fields.data,
				thetas: tt.fields.thetas,
			}
			if got := lr.calcGradient(tt.args.residuals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcGradient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinearRegression_calcError(t *testing.T) {
	type fields struct {
		config LinearRegressionConfig
		data   LinearRegressionData
		thetas Vec
	}
	type args struct {
		p Vec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		{
			name: "Calculate error.",
			fields: fields{
				LinearRegressionConfig{
					learningRate:   0.001,
					printUpdates:   false,
					printFrequency: 100,
					epochs:         100000,
				},
				LinearRegressionData{
					x: Vec{1, 2, 3, 4}.AddOnes(),
					y: Vec{3, 4, 5, 6},
					l: 4,
				},
				Vec{1, 1},
			},
			args: args{
				p: Vec{1, 2, 3, 4},
			},
			want: 16.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &LinearRegression{
				config: tt.fields.config,
				data:   tt.fields.data,
				thetas: tt.fields.thetas,
			}
			if got := lr.calcError(tt.args.p); got != tt.want {
				t.Errorf("calcError() = %v, want %v", got, tt.want)
			}
		})
	}
}
