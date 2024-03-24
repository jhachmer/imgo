package filter

import (
	"reflect"
	"testing"
)

func TestCreateGaussKernel1D(t *testing.T) {
	type args struct {
		sigma float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateGaussKernel1D(tt.args.sigma); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateGaussKernel1D() = %v, want %v", got, tt.want)
			}
		})
	}
}
