package leetcode

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: struct {
				s       string
				numRows int
			}{s: "PAYPALISHIRING", numRows: 3},
			want: "PAHNAPLSIIGYIR",
		},
		{name: "2",
			args: struct {
				s       string
				numRows int
			}{s: "PAYPALISHIRING", numRows: 4},
			want: "PINALSIGYAHRPI",
		},
		{name: "3",
			args: struct {
				s       string
				numRows int
			}{s: "A", numRows: 1},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
