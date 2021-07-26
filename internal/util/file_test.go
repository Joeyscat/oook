package util

import (
	"testing"
)

func TestDirExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test ok",
			args{"/etc"},
			true,
		},
		{
			"test ok",
			args{"/etc/x"},
			false,
		},
		{
			"test ok",
			args{"/etc/hosts"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirExists(tt.args.filename); got != tt.want {
				t.Errorf("DirExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test ok",
			args{"/etc"},
			false,
		},
		{
			"test ok",
			args{"/etc/x"},
			false,
		},
		{
			"test ok",
			args{"/etc/hosts"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExists(tt.args.filename); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
