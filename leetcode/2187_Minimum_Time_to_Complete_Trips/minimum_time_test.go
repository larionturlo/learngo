package main

import (
	"testing"
)

func Test_minimumTime(t *testing.T) {
	type args struct {
		time       []int
		totalTrips int
	}

	tests := []struct {
		name string
		arg  args
		want int64
	}{
		// {
		// 	"#0",
		// 	args{
		// 		time:       []int{1, 2, 3},
		// 		totalTrips: 5,
		// 	},
		// 	3,
		// },
		{
			"#0",
			args{
				time:       []int{5, 10, 10},
				totalTrips: 9,
			},
			25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTime(tt.arg.time, tt.arg.totalTrips); got != tt.want {
				t.Errorf("#%s  = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
