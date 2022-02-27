package algo

import "testing"

func Test_bynarySearch(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			"Test",
			args{[]int{1, 2, 3, 4, 5}, 2},
			1, true,
		},
		{
			"Test",
			args{[]int{1, 2, 3, 4, 5}, 4},
			3, true,
		},
		{
			"Test",
			args{[]int{1, 2, 3, 5, 6}, 4},
			0, false,
		},
		{
			"Test",
			args{[]int{1, 3, 4, 5, 6}, 2},
			0, false,
		},
		{
			"Test",
			args{[]int{1, 2, 3, 4, 5}, 6},
			0, false,
		},
		{
			"Test",
			args{[]int{1, 2, 3, 4, 5}, 0},
			0, false,
		},
		{
			"Test",
			args{[]int{1, 2, 3, 4, 5}, 5},
			4, true,
		},
		{
			"Test",
			args{[]int{1, 2, 3, 4, 5}, 1},
			0, true,
		},
		{
			"Test",
			args{[]int{1, 2, 3, 4, 5}, 3},
			2, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := bynarySearch(tt.args.arr, tt.args.n)
			if got != tt.want {
				t.Errorf("bynarySearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("bynarySearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
