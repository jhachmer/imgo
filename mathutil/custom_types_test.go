package mathutil

import (
	"reflect"
	"testing"
)

func TestCalcMagnitude(t *testing.T) {
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
