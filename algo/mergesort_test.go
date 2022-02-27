package algo

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name  string
		args  args
		wantR []int
	}{
		{
			"test",
			args{
				[]int{0, 2, 3, 7, 10},
				[]int{1, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10, 10, 11, 12},
		},
		{
			"test",
			args{
				[]int{0, 2},
				[]int{1, 3, 6},
			},
			[]int{0, 1, 2, 3, 6},
		},
		{
			"test",
			args{
				[]int{0},
				[]int{0},
			},
			[]int{0, 0},
		},
		{
			"test",
			args{
				[]int{},
				[]int{0, 1, 5},
			},
			[]int{0, 1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := merge(tt.args.a, tt.args.b); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("merge() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test",
			args{
				[]int{0, 5, 6, 2, 19, -1, 2, 3, 0, 4, 5, 8},
			},
			[]int{-1, 0, 0, 2, 2, 3, 4, 5, 5, 6, 8, 19},
		},
		{
			"test",
			args{
				[]int{0, 1, 2, 6, 18, 19, -20, -45, -23, 25, 56, 19, 81, 123, 122},
			},
			[]int{-45, -23, -20, 0, 1, 2, 6, 18, 19, 19, 25, 56, 81, 122, 123},
		},
		{
			"test",
			args{
				[]int{4, 3, 2, 1},
			},
			[]int{1, 2, 3, 4},
		},
		{
			"test",
			args{
				[]int{1},
			},
			[]int{1},
		},
		{
			"test",
			args{
				[]int{},
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
