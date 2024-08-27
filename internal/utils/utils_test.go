package utils

import (
	"testing"
)

func TestSizeToString(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0 bytes = 0B",
			args: args{n: 0},
			want: "0B",
		},
		{
			name: "23456 bytes are 22.9 KB",
			args: args{n: 23456},
			want: "22.9KB",
		},
		{
			name: "123456789 bytes are 117.7MB",
			args: args{n: 123456789},
			want: "117.7MB",
		},
		{
			name: "258997 bytes are 252.9KB",
			args: args{n: 258997},
			want: "252.9KB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SizeToString(tt.args.n); got != tt.want {
				t.Errorf("SizeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
