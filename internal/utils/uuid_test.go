package utils

import (
	"strings"
	"testing"

	uuid2 "github.com/google/uuid"
)

func TestIsValidUUID(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is valid UUID",
			args: args{uuid: uuid2.New().String()},
			want: true,
		},
		{
			name: "is not a valid UUID",
			args: args{uuid: "123123123123"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidUUID(tt.args.uuid); got != tt.want {
				t.Errorf("IsValidUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilenameWithUUID(t *testing.T) {
	type args struct {
		filename string
	}
	type result struct {
		filename string
		isValid  bool
	}
	tests := []struct {
		name string
		args args
		want result
	}{
		{
			name: "image.jpg",
			args: args{filename: "image.jpg"},
			want: result{
				filename: "image.jpg",
				isValid:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilenameWithUUID(tt.args.filename)
			sep := strings.SplitAfterN(got, "_", 2)
			if len(sep) != 2 {
				t.Errorf("FilenameWithUUID() did not return string with separator")
			}
			if IsValidUUID(strings.TrimSuffix(sep[0], "_")) != tt.want.isValid {
				t.Errorf("FilenameWithUUID() UUID part = %v, want isValid?: %v", sep[0], tt.want.isValid)
			}
			if sep[1] != tt.want.filename {
				t.Errorf("FilenameWithUUID() = %v, want %v", sep[1], tt.want.filename)
			}
		})
	}
}
