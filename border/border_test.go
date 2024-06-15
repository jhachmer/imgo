package border

import "testing"

func TestBorderDetection(t *testing.T) {
	type args struct {
		u    int
		v    int
		i    int
		j    int
		xMax int
		yMax int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{name: "not at border (2,2)", args: args{
			u:    2,
			v:    2,
			i:    0,
			j:    0,
			xMax: 10,
			yMax: 10,
		},
			want:  2,
			want1: 2,
		},
		{name: "right on border (0,0)", args: args{
			u:    1,
			v:    1,
			i:    -1,
			j:    -1,
			xMax: 10,
			yMax: 10,
		},
			want:  0,
			want1: 0,
		},
		{name: "behind border (0,0)", args: args{
			u:    0,
			v:    0,
			i:    -1,
			j:    -1,
			xMax: 10,
			yMax: 10,
		},
			want:  1,
			want1: 1,
		},
		{name: "behind border (x)", args: args{
			u:    0,
			v:    5,
			i:    -1,
			j:    -1,
			xMax: 10,
			yMax: 10,
		},
			want:  1,
			want1: 4,
		},
		{name: "behind border (y)", args: args{
			u:    5,
			v:    0,
			i:    -1,
			j:    -1,
			xMax: 10,
			yMax: 10,
		},
			want:  4,
			want1: 1,
		},
		{name: "behind border max", args: args{
			u:    10,
			v:    10,
			i:    1,
			j:    1,
			xMax: 10,
			yMax: 10,
		},
			want:  9,
			want1: 9,
		},
		{name: "behind border max (x)", args: args{
			u:    10,
			v:    5,
			i:    1,
			j:    1,
			xMax: 10,
			yMax: 10,
		},
			want:  9,
			want1: 6,
		},
		{name: "behind border max (y)", args: args{
			u:    5,
			v:    10,
			i:    1,
			j:    1,
			xMax: 10,
			yMax: 10,
		},
			want:  6,
			want1: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Detection(tt.args.u, tt.args.v, tt.args.i, tt.args.j, tt.args.xMax, tt.args.yMax)
			if got != tt.want {
				t.Errorf("Detection() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Detection() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
