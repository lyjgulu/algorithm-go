package leetcode

import (
	"algorithm-go/structures"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: struct {
				l1 *ListNode
				l2 *ListNode
			}{l1: structures.Ints2List([]int{}), l2: structures.Ints2List([]int{})},
			want: structures.Ints2List([]int{}),
		},
		{
			name: "2",
			args: struct {
				l1 *ListNode
				l2 *ListNode
			}{l1: structures.Ints2List([]int{1}), l2: structures.Ints2List([]int{1})},
			want: structures.Ints2List([]int{2}),
		},
		{
			name: "3",
			args: struct {
				l1 *ListNode
				l2 *ListNode
			}{l1: structures.Ints2List([]int{9, 9, 9, 9}), l2: structures.Ints2List([]int{1})},
			want: structures.Ints2List([]int{0, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
