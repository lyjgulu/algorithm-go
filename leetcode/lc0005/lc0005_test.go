package leetcode

import "testing"

func Test_longestPalindrome3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: struct {
				s string
			}{s: "babad"},
			want: "bab",
		},
		{
			name: "2",
			args: struct {
				s string
			}{s: "cbbd"},
			want: "bb",
		},
		{
			name: "3",
			args: struct {
				s string
			}{s: "ac"},
			want: "a",
		},
		{
			name: "3",
			args: struct {
				s string
			}{s: "aa"},
			want: "aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome3(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome3() = %v, want %v", got, tt.want)
			}
		})
	}
}
