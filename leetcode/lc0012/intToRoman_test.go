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
		{name: testing.CoverMode(), args: args{num: 3}, want: "III"},
		{name: testing.CoverMode(), args: args{num: 4}, want: "IV"},
		{name: testing.CoverMode(), args: args{num: 9}, want: "IX"},
		{name: testing.CoverMode(), args: args{num: 58}, want: "LVIII"},
		{name: testing.CoverMode(), args: args{num: 1994}, want: "MCMXCIV"},
		{name: testing.CoverMode(), args: args{num: 123}, want: "CXXIII"},
		{name: testing.CoverMode(), args: args{num: 120}, want: "CXX"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intTORoman(tt.args.num); got != tt.want {
				t.Errorf("intTORoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
