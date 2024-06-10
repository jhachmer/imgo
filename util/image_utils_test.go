package util

import (
	"image"
	"reflect"
	"testing"
)

func testImage(pix []uint8, stride int, rect image.Rectangle) *image.Gray {
	return &image.Gray{Pix: pix, Stride: stride, Rect: rect}
}

func TestImageToSlice(t *testing.T) {
	type args struct {
		img *image.Gray
	}
	tests := []struct {
		name string
		args args
		want [][]uint8
	}{
		{
			name: "2x2",
			args: args{img: testImage([]uint8{1, 1, 1, 1}, 2, image.Rect(0, 0, 2, 2))},
			want: [][]uint8{{1, 1}, {1, 1}},
		},
		{
			name: "4x4",
			args: args{img: testImage([]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 4, image.Rect(0, 0, 4, 4))},
			want: [][]uint8{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ImageToSlice(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImageToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
