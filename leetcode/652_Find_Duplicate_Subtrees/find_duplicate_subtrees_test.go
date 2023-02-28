package main

import (
	"reflect"
	"testing"
)

func Test_find_duplicate_subtrees(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		// {
		// 	"simple ",
		// 	args{
		// 		root: &TreeNode{
		// 			Val: 2,
		// 			Left: &TreeNode{
		// 				Val: 1,
		// 			},
		// 			Right: &TreeNode{
		// 				Val: 1,
		// 			},
		// 		},
		// 	},
		// 	[]*TreeNode{&TreeNode{
		// 		Val: 1,
		// 	}},
		// },

		// {
		// 	"complex ",
		// 	args{
		// 		root: &TreeNode{
		// 			Val: 1,
		// 			Left: &TreeNode{
		// 				Val: 2,
		// 				Left: &TreeNode{
		// 					Val: 4,
		// 				},
		// 			},
		// 			Right: &TreeNode{
		// 				Val: 3,
		// 				Left: &TreeNode{
		// 					Val: 2,
		// 					Left: &TreeNode{
		// 						Val: 4,
		// 					},
		// 				},
		// 				Right: &TreeNode{
		// 					Val: 4,
		// 				},
		// 			},
		// 		},
		// 	},
		// 	[]*TreeNode{
		// 		&TreeNode{
		// 			Val: 2,
		// 			Left: &TreeNode{
		// 				Val: 4,
		// 			},
		// 		},
		// 		&TreeNode{
		// 			Val: 4,
		// 		},
		// 	},
		// },

		{
			"ziro ",
			args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
					},
					Right: &TreeNode{
						Val: 0,
						Right: &TreeNode{
							Val: 0,
							Left: &TreeNode{
								Val: 0,
							},
							Right: &TreeNode{
								Val: 0,
							},
						},
					},
				},
			},
			[]*TreeNode{
				&TreeNode{
					Val: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicateSubtrees(tt.args.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}