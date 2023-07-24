package leetcode

import "testing"

func Test_intTORoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intTORoman(tt.args.num); got != tt.want {
				t.Errorf("intTORoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
