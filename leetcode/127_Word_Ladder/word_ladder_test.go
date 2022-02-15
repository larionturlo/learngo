package main

import "testing"

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				"hot",
				"dog",
				[]string{"hot", "dog", "dot"},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				"hit",
				"cog",
				[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 5,
		},
		{
			name: "3",
			args: args{
				"hit",
				"cog",
				[]string{"hot", "dot", "dog", "lot", "log"},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				"a",
				"c",
				[]string{"a", "b", "c"},
			},
			want: 2,
		},
		{
			name: "5",
			args: args{
				"a",
				"b",
				[]string{"a", "b", "c"},
			},
			want: 2,
		},
		/*
					"hot"
			"dog"
			["hot","dot","dog"]*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
