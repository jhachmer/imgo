package util

import (
	"reflect"
	"testing"

	"github.com/jhachmer/imgo/model"
)

func ComplexSlice() [][]model.Complex {
	return [][]model.Complex{
		{
			{Re: 1, Im: 0}, {Re: 2, Im: 0},
		},
		{
			{Re: 3, Im: 0}, {Re: 4, Im: 0},
		},
	}
}

func TestGenerateComplexSlice(t *testing.T) {
	type args struct {
		pixelLine [][]uint8
	}
	tests := []struct {
		name string
		args args
		want [][]model.Complex
	}{
		{
			name: "2",
			args: args{[][]uint8{{1, 2}, {3, 4}}},
			want: ComplexSlice(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateComplexSlice(tt.args.pixelLine); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateComplexSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}