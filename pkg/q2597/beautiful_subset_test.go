package q2597

import "testing"

func Test_beautifulSubsets(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//输入：nums = [2,4,6], k = 2
		//输出：4
		{"1", args{[]int{2, 4, 6}, 2}, 4},
		{"2", args{[]int{1}, 1}, 1},
		{"3", args{[]int{4, 2, 5, 9, 10, 3}, 1}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautifulSubsets(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("beautifulSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
