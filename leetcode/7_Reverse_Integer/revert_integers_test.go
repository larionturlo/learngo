package main

import (
	"testing"
)

func Test_reverse_integer(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			"#0",
			123,
			321,
		},
		{
			"#1",
			-123,
			-321,
		}, {
			"#2",
			120,
			21,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse_integer(tt.arg); got != tt.want {
				t.Errorf("#%s reverse_integer(%d) = %v, want %v", tt.name, tt.arg, got, tt.want)
			}
		})
	}
}
func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			"#0",
			123,
			321,
		},
		{
			"#1",
			-123,
			-321,
		}, {
			"#2",
			120,
			21,
		}, {
			"#3",
			1534236469,
			0,
		}, {
			"#4",
			-2147483648,
			0,
		},}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.arg); got != tt.want {
				t.Errorf("#%s reverse(%d) = %v, want %v", tt.name, tt.arg, got, tt.want)
			}
		})
	}
}