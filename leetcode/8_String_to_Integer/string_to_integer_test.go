package main

import (
	"testing"
)

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			"#0",
			"123",
			123,
		},
		{
			"#1",
			"-123",
			-123,
		}, {
			"#2",
			"0120",
			120,
		},
		{
			"#3",
			"  -123",
			-123,
		},
		{
			"#4",
			"-123 were",
			-123,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.arg); got != tt.want {
				t.Errorf("#%s reverse_integer(\"%s\") = %v, want %v", tt.name, tt.arg, got, tt.want)
			}
		})
	}
}
