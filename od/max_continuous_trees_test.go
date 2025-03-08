package od

import "testing"

func TestMaxContinuousTrees(t *testing.T) {
	type args struct {
		n           int
		m           int
		deadIndices []int
		k           int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				n:           5,
				m:           2,
				deadIndices: []int{2, 4},
				k:           1,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				n:           10,
				m:           3,
				deadIndices: []int{2, 4, 7},
				k:           1,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxContinuousTrees(tt.args.n, tt.args.m, tt.args.deadIndices, tt.args.k); got != tt.want {
				t.Errorf("MaxContinuousTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
