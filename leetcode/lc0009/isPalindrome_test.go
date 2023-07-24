package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: testing.CoverMode(),
			args: args{x: 121},
			want: true,
		},
		{
			name: testing.CoverMode(),
			args: args{x: -121},
			want: false,
		},
		{
			name: testing.CoverMode(),
			args: args{x: 10},
			want: false,
		},
		{
			name: testing.CoverMode(),
			args: args{x: 1534236469},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
