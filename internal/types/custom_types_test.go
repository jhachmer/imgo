package types

import (
	"github.com/google/go-cmp/cmp"
	"math"
	"math/cmplx"
	"reflect"
	"testing"
)

const tolerance = .0001

func TestGradient2D_CalcMagnitude(t *testing.T) {
	tests := []struct {
		name string
		grad Gradient2D
		want int
	}{
		{name: "First (1,1)", grad: Gradient2D{X: 1.0, Y: 1.0}, want: 2},
		{name: "Second (10,10)", grad: Gradient2D{X: 10, Y: 10}, want: 15},
		{name: "Third (100,100)", grad: Gradient2D{X: 100, Y: 100}, want: 142},
		{name: "Fourth (1000,1000)", grad: Gradient2D{X: 1000, Y: 1000}, want: 255},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.grad.CalcMagnitude(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcMagnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplexRect(t *testing.T) {
	opt := cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		return delta/mean < 0.00001
	})
	r1, phi1 := cmplx.Polar(1 + 1i)
	r2, phi2 := cmplx.Polar(2 + 2i)
	r3, phi3 := cmplx.Polar(3 + 3i)

	type args struct {
		mag float64
		pha float64
	}
	tests := []struct {
		name string
		args args
		want *Complex
	}{
		{
			name: "Test #1 1+1i",
			args: args{mag: r1, pha: phi1},
			want: &Complex{
				Re: 1,
				Im: 1,
			},
		},
		{
			name: "Test #2 2+2i",
			args: args{mag: r2, pha: phi2},
			want: &Complex{
				Re: 2,
				Im: 2,
			},
		},
		{
			name: "Test #3 3+3i",
			args: args{mag: r3, pha: phi3},
			want: &Complex{
				Re: 3,
				Im: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComplexRect(tt.args.mag, tt.args.pha); !cmp.Equal(got.Re, tt.want.Re, opt) || !cmp.Equal(got.Im, tt.want.Im, opt) {
				t.Errorf("ComplexRect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex_Abs(t *testing.T) {
	type fields struct {
		Re float64
		Im float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Test #1 in: (1, 1) want: sqrt(2)",
			fields: fields{Re: 1, Im: 1},
			want:   math.Sqrt2,
		},
		{
			name:   "Test #2 in: (2, 2) want: sqrt(2^2+2^2)",
			fields: fields{Re: 2, Im: 2},
			want:   math.Hypot(2, 2),
		},
		{
			name:   "Test #3 in: (16, 65) want: sqrt(16^2+65^2)",
			fields: fields{Re: 16, Im: 65},
			want:   math.Hypot(16, 65),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Complex{
				Re: tt.fields.Re,
				Im: tt.fields.Im,
			}
			if got := c.Abs(); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex_Phase(t *testing.T) {
	type fields struct {
		Re float64
		Im float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test #1 in: (1, 1) want: arctan2(1/1)",
			fields: fields{
				Re: 1,
				Im: 1,
			},
			want: math.Atan2(1, 1),
		},
		{
			name: "Test #2 in: (2, 2) want: arcta2(2/2)",
			fields: fields{
				Re: 2,
				Im: 2,
			},
			want: math.Atan2(2, 2),
		},
		{
			name: "Test #3 in: (15, 23) want: arctan2(23/15)",
			fields: fields{
				Re: 15,
				Im: 23,
			},
			want: math.Atan2(23, 15),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Complex{
				Re: tt.fields.Re,
				Im: tt.fields.Im,
			}
			if got := c.Phase(); got != tt.want {
				t.Errorf("Phase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewComplex(t *testing.T) {
	type args struct {
		re float64
		im float64
	}
	tests := []struct {
		name string
		args args
		want *Complex
	}{
		{
			name: "Test #1 (1,1)",
			args: args{
				re: 1,
				im: 1,
			},
			want: &Complex{
				Re: 1,
				Im: 1,
			},
		},
		{
			name: "Test #2 (5,5)",
			args: args{
				re: 5,
				im: 5,
			},
			want: &Complex{
				Re: 5,
				Im: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComplex(tt.args.re, tt.args.im); !cmp.Equal(got, tt.want) {
				t.Errorf("NewComplex() = %v, want %v", got, tt.want)
			}
		})
	}
}
