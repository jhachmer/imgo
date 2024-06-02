package mathutil

import "testing"

func TestAbs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0 is 0", args: args{0}, want: 0},
		{name: "1 is 1", args: args{1}, want: 1},
		{name: "-1 is 1", args: args{-1}, want: 1},
		{name: "50 is 50", args: args{50}, want: 50},
		{name: "-50 is 50", args: args{-50}, want: 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
