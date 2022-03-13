package heap

import (
	"reflect"
	"testing"
)

func TestHeap_Heapify(t *testing.T) {
	type fields struct {
		heapType HeapType
	}
	type args struct {
		arr []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			"MinHeap",
			fields{
				minHeap,
			},
			args{
				[]int{5, 4, 4, 3, 3, 2, 2, 1},
			},
			[]int{1, 2, 3, 2, 5, 4, 4, 3},
		},
		{
			"MaxHeap",
			fields{
				maxHeap,
			},
			args{
				[]int{1, 2, 3, 2, 5, 4, 4, 3},
			},
			[]int{5, 4, 3, 4, 1, 2, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHeap(tt.fields.heapType)
			h.Heapify(tt.args.arr)
			if got := h.AsSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_Insert(t *testing.T) {
	type fields struct {
		heap     []int
		heapType HeapType
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			"minHeap",
			fields{
				[]int{},
				minHeap,
			},
			args{5},
			[]int{5},
		},
		{
			"minHeap",
			fields{
				[]int{4, 5},
				minHeap,
			},
			args{3},
			[]int{3, 4, 5},
		},
		{
			"maxHeap",
			fields{
				[]int{4, 3},
				maxHeap,
			},
			args{5},
			[]int{5, 4, 3},
		},
		{
			"maxHeap",
			fields{
				[]int{5, 4, 3, 1, 2, 2, 3},
				maxHeap,
			},
			args{4},
			[]int{5, 4, 3, 4, 2, 2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				heap:     tt.fields.heap,
				heapType: tt.fields.heapType,
			}
			h.Insert(tt.args.n)
			if got := h.AsSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_Delete(t *testing.T) {
	type fields struct {
		heap     []int
		heapType HeapType
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
		want2  int
	}{
		{
			"minHeap",
			fields{
				[]int{},
				minHeap,
			},
			args{0},
			[]int{},
			0,
		},
		{
			"minHeap",
			fields{
				[]int{3, 4, 5},
				minHeap,
			},
			args{0},
			[]int{4, 5},
			3,
		},
		{
			"maxHeap",
			fields{
				[]int{5, 4, 3},
				maxHeap,
			},
			args{0},
			[]int{4, 3},
			5,
		},
		{
			"maxHeap",
			fields{
				[]int{5, 4, 3, 4, 2, 2, 3, 1},
				maxHeap,
			},
			args{3},
			[]int{5, 4, 3, 3, 2, 2, 1},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				heap:     tt.fields.heap,
				heapType: tt.fields.heapType,
			}
			got2 := h.Extract(tt.args.i)
			if got := h.AsSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heap after Extract() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("result Extract() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
