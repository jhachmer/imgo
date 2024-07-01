package filter

import (
	"reflect"
	"testing"

	"github.com/jhachmer/imgo/pkg/kernel"
)

func TestCreateGaussKernel1D(t *testing.T) {
	type args struct {
		sigma float64
	}
	tests := []struct {
		name string
		args args
		want *kernel.Kernel1D
	}{
		{name: "sigma 1", args: args{sigma: 1.0}, want: &kernel.Kernel1D{
			Size:   248,
			Values: []int{1, 13, 60, 100, 60, 13, 1},
			Len:    7,
		}},
		{name: "sigma 5", args: args{sigma: 5.0}, want: &kernel.Kernel1D{
			Size:   1234,
			Values: []int{1, 1, 3, 5, 8, 13, 19, 27, 37, 48, 60, 72, 83, 92, 98, 100, 98, 92, 83, 72, 60, 48, 37, 27, 19, 13, 8, 5, 3, 1, 1},
			Len:    31,
		}},
		{name: "sigma3", args: args{sigma: 3.0}, want: &kernel.Kernel1D{
			Size:   742,
			Values: []int{1, 2, 6, 13, 24, 41, 60, 80, 94, 100, 94, 80, 60, 41, 24, 13, 6, 2, 1},
			Len:    19,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateGaussKernel1D(tt.args.sigma); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateGaussKernel1D() = %v, want %v", got, tt.want)
			}
		})
	}
}
